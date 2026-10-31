// Command smoke calls every operation of the SDK against a real Affinity API in Test mode.
//
//	AFFINITY_API_KEY=sk_test_... [AFFINITY_BASE_URL=https://...] [AFFINITY_SMOKE_WEBHOOK_URL=https://...] \
//	  go run ./tests/smoke
//
// It refuses a Live key. Use a platform Test key: it creates a practice and a Test order (sent to
// the API's test pharmacy) and cancels the order. With AFFINITY_SMOKE_WEBHOOK_URL it also creates,
// updates and removes a webhook endpoint. Feedback is sent only as an invalid request, so no one
// is notified. The Test practice it creates stays: the API has no way to remove a practice.
// It prints PASS or FAIL per check and exits 1 when any check fails.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	affinity "github.com/affinity-health/affinity-go"
	"github.com/affinity-health/affinity-go/client"
	"github.com/affinity-health/affinity-go/option"
	"github.com/google/uuid"
)

type result struct {
	name string
	ok   bool
	note string
}

var results []result

func record(name string, err error, note string) {
	if err != nil {
		results = append(results, result{name, false, describe(err)})
		return
	}
	results = append(results, result{name, true, note})
}

func describe(err error) string {
	text := err.Error()
	if len(text) > 300 {
		text = text[:300]
	}
	return fmt.Sprintf("%T: %s", err, text)
}

// refused records a check that passes only when err is an E and inspect (if any) finds no problem.
func refused[E error](name string, err error, inspect func(E) string) {
	var target E
	switch {
	case err == nil:
		results = append(results, result{name, false, "succeeded"})
	case !errors.As(err, &target):
		results = append(results, result{name, false, "got " + describe(err)})
	default:
		problem := ""
		if inspect != nil {
			problem = inspect(target)
		}
		results = append(results, result{name, problem == "", firstNonEmpty(problem, describe(err))})
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func key() string { return uuid.NewString() }

func main() {
	apiKey := os.Getenv("AFFINITY_API_KEY")
	if !strings.HasPrefix(apiKey, "sk_test_") {
		fmt.Fprintln(os.Stderr, "Set AFFINITY_API_KEY to a Test key (sk_test_…).")
		os.Exit(1)
	}
	baseURL := os.Getenv("AFFINITY_BASE_URL")
	webhookURL := strings.TrimRight(os.Getenv("AFFINITY_SMOKE_WEBHOOK_URL"), "/")
	options := []option.RequestOption{option.WithToken(apiKey), option.WithoutRetries()}
	if baseURL != "" {
		options = append(options, option.WithBaseURL(baseURL))
	}
	c := client.NewClient(options...)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	var orderIDs, endpointIDs []string
	func() {
		defer func() {
			// Cancel every order this run created and remove every endpoint, even after a failure.
			for _, id := range orderIDs {
				_, _ = c.Orders.Cancel(ctx, &affinity.CancelOrdersRequest{OrderID: id, IdempotencyKey: key()})
			}
			for _, id := range endpointIDs {
				_, _ = c.Webhooks.Remove(ctx, &affinity.RemoveWebhooksRequest{WebhookEndpointID: id})
			}
		}()
		if err := run(ctx, c, baseURL, webhookURL, &orderIDs, &endpointIDs); err != nil {
			results = append(results, result{"stopped", false, err.Error()})
		}
	}()

	failed := 0
	for _, r := range results {
		status := "PASS"
		if !r.ok {
			status = "FAIL"
			failed++
		}
		line := status + " " + r.name
		if r.note != "" {
			line += " — " + r.note
		}
		fmt.Println(line)
	}
	fmt.Printf("%d of %d checks passed.\n", len(results)-failed, len(results))
	if failed > 0 {
		os.Exit(1)
	}
}

func run(ctx context.Context, c *client.Client, baseURL, webhookURL string, orderIDs, endpointIDs *[]string) error {
	label := "SDK smoke " + time.Now().UTC().Format(time.RFC3339)

	account, err := c.Account.Get(ctx)
	record("account.get", err, func() string {
		if account == nil {
			return ""
		}
		return string(account.Type) + " " + string(account.Mode)
	}())
	if err != nil {
		return errors.New("account.get failed")
	}
	if account.Mode != affinity.AccountModeTest {
		return errors.New("this smoke test runs only with a Test key")
	}
	if account.Type != affinity.AccountTypePlatform {
		return errors.New("use a platform Test key: it creates a practice")
	}

	badOptions := []option.RequestOption{option.WithToken("sk_test_not_a_key")}
	if baseURL != "" {
		badOptions = append(badOptions, option.WithBaseURL(baseURL))
	}
	_, err = client.NewClient(badOptions...).Account.Get(ctx)
	refused[*affinity.UnauthorizedError]("a bad key raises UnauthorizedError", err, nil)

	_, err = c.Feedback.Create(ctx, &affinity.CreateFeedbackRequest{IdempotencyKey: key(), Message: ""})
	refused[*affinity.BadRequestError]("feedback.create without a message raises BadRequestError", err, nil)

	// Practices.
	practice, err := c.Practices.Create(ctx, &affinity.CreatePracticesRequest{
		IdempotencyKey: key(),
		Name:           label,
		Phone:          "+15125550100",
		Address:        &affinity.AddressInput{Line1: "100 Congress Ave", City: "Austin", State: "TX", PostalCode: "78701"},
	})
	record("practices.create", err, idOf(practice))
	if err != nil {
		return errors.New("no practice")
	}
	got, err := c.Practices.Get(ctx, &affinity.GetPracticesRequest{PracticeID: practice.ID})
	record("practices.get", err, func() string {
		if got == nil {
			return ""
		}
		return got.Name
	}())
	updated, err := c.Practices.Update(ctx, &affinity.UpdatePracticesRequest{
		PracticeID:     practice.ID,
		IdempotencyKey: key(),
		Name:           affinity.String(label + " renamed"),
	})
	if err == nil && updated.Name != label+" renamed" {
		err = fmt.Errorf("name is %q", updated.Name)
	}
	record("practices.update", err, label+" renamed")
	{
		page, err := c.Practices.List(ctx, &affinity.ListPracticesRequest{Limit: affinity.Int(1)})
		found, n := false, 0
		if err == nil {
			it := page.Iterator()
			for it.Next(ctx) {
				n++
				found = found || it.Current().ID == practice.ID
			}
			err = it.Err()
		}
		if err == nil && !found {
			err = errors.New("the new practice is not listed")
		}
		record("practices.list (every page, 1 per page)", err, fmt.Sprintf("%d practices", n))
	}

	// Medications.
	{
		page, err := c.Medications.List(ctx, &affinity.ListMedicationsRequest{Limit: affinity.Int(5)})
		note := ""
		if err == nil {
			note = fmt.Sprintf("%d on the first page", len(page.Results))
		}
		record("medications.list (one page)", err, note)
	}
	var medications []*affinity.Medication
	{
		page, err := c.Medications.List(ctx, &affinity.ListMedicationsRequest{Limit: affinity.Int(25)})
		if err == nil {
			it := page.Iterator()
			for it.Next(ctx) {
				medications = append(medications, it.Current())
			}
			err = it.Err()
		}
		seen := map[string]bool{}
		for _, m := range medications {
			if seen[m.ID] && err == nil {
				err = fmt.Errorf("medication %s repeats across pages", m.ID)
			}
			seen[m.ID] = true
		}
		record("medications.list (every page, 25 per page)", err, fmt.Sprintf("%d medications", len(medications)))
	}
	var medication *affinity.Medication
	for _, m := range medications {
		if !m.ColdChainRequired && len(m.ShippingOptions) > 0 && len(m.Requirements) == 0 {
			medication = m
			break
		}
	}
	for _, m := range medications {
		if medication == nil && len(m.ShippingOptions) > 0 {
			medication = m
		}
	}
	if medication == nil {
		return errors.New("no orderable medication in Test mode")
	}
	{
		m, err := c.Medications.Get(ctx, &affinity.GetMedicationsRequest{MedicationID: medication.ID})
		note := ""
		if err == nil {
			note = m.Name + " (" + m.Pharmacy.Name + ")"
		}
		record("medications.get", err, note)
		variant := affinity.GetMedicationsRequestExpandItemVariants
		m, err = c.Medications.Get(ctx, &affinity.GetMedicationsRequest{
			MedicationID: medication.ID,
			Expand:       []*affinity.GetMedicationsRequestExpandItem{&variant},
		})
		if err == nil && m.Variants == nil {
			err = errors.New("no variants field in the response")
		}
		if err == nil {
			note = fmt.Sprintf("%d variants", len(m.Variants))
		}
		record("medications.get expand=variants", err, note)
		page, err := c.Medications.List(ctx, &affinity.ListMedicationsRequest{Limit: affinity.Int(5), PracticeID: affinity.String(practice.ID)})
		if err == nil {
			note = fmt.Sprintf("%d", len(page.Results))
		}
		record("medications.list for a practice", err, note)
	}

	// Orders.
	orderInput := func(quantity int) *affinity.OrderInput {
		prescription := &affinity.PrescriptionInput{
			MedicationID: medication.ID,
			Quantity:     quantity,
			Refills:      0,
			Directions:   "SDK smoke test. TEST ONLY. DO NOT FILL.",
		}
		for _, r := range medication.Requirements {
			if r == "compoundingReason" && len(medication.CompoundingReasons) > 0 {
				prescription.CompoundingReason = affinity.String(medication.CompoundingReasons[0])
			}
		}
		return &affinity.OrderInput{
			IdempotencyKey:  key(),
			ExternalOrderID: affinity.String("sdk-smoke-" + key()[:8]),
			Patient: &affinity.PatientInput{
				ExternalID:  affinity.String("sdk-smoke-" + key()),
				FirstName:   "Test",
				LastName:    "Smoke",
				DateOfBirth: "1988-04-02",
				Sex:         affinity.PatientInputSexFemale,
				Email:       "sdk-smoke-" + key()[:8] + "@example.test",
				Phone:       "+15125550101",
				Allergies:   []*affinity.PatientInputAllergiesItem{},
				Address:     &affinity.AddressInput{Line1: "200 Congress Ave", City: "Austin", State: "TX", PostalCode: "78701"},
			},
			PracticeID:           affinity.String(practice.ID),
			Prescriber:           &affinity.PrescriberInput{Npi: "1234567893", FirstName: "Sam", LastName: "Lee"},
			Prescriptions:        []*affinity.PrescriptionInput{prescription},
			Shipping:             &affinity.OrderInputShipping{OptionID: medication.ShippingOptions[0].ID},
			SignatureAttestation: true,
		}
	}
	quantity := 1
	if len(medication.Quantities) > 0 {
		quantity = medication.Quantities[0]
	}
	input := orderInput(quantity)
	order, err := c.Orders.Create(ctx, input)
	record("orders.create", err, func() string {
		if order == nil {
			return ""
		}
		return order.ID + " " + string(order.Status)
	}())
	if order != nil {
		*orderIDs = append(*orderIDs, order.ID)
	}
	{
		again, err := c.Orders.Create(ctx, input)
		if again != nil && (order == nil || again.ID != order.ID) {
			*orderIDs = append(*orderIDs, again.ID)
			if err == nil {
				err = fmt.Errorf("created a second order %s", again.ID)
			}
		}
		record("orders.create again with the same key returns the same order", err, idOf(again))
	}
	{
		highest := 0
		for _, q := range medication.Quantities {
			highest = max(highest, q)
		}
		unexpected, err := c.Orders.Create(ctx, orderInput(highest+1))
		if unexpected != nil {
			*orderIDs = append(*orderIDs, unexpected.ID)
		}
		refused("an unaccepted quantity raises UnprocessableEntityError with the field", err,
			func(e *affinity.UnprocessableEntityError) string {
				for _, issue := range e.GetBody().GetIssues() {
					if issue.Field == "prescriptions[0].quantity" {
						return ""
					}
				}
				return fmt.Sprintf("issues %+v", e.GetBody().GetIssues())
			})
	}
	if order != nil {
		got, err := c.Orders.Get(ctx, &affinity.GetOrdersRequest{OrderID: order.ID})
		note := ""
		if err == nil {
			note = fmt.Sprintf("%s, %d prescription", got.Status, len(got.Prescriptions))
		}
		record("orders.get", err, note)

		page, err := c.Orders.List(ctx, &affinity.ListOrdersRequest{ExternalOrderID: input.ExternalOrderID})
		if err == nil && (len(page.Results) != 1 || page.Results[0].ID != order.ID) {
			err = fmt.Errorf("found %d orders", len(page.Results))
		}
		record("orders.list by externalOrderId", err, "")

		hourAgo := time.Now().Add(-time.Hour).Truncate(time.Millisecond)
		page, err = c.Orders.List(ctx, &affinity.ListOrdersRequest{UpdatedAfter: &hourAgo, Limit: affinity.Int(5)})
		if err == nil {
			note = fmt.Sprintf("%d", len(page.Results))
		}
		record("orders.list updatedAfter (milliseconds)", err, note)

		nanos, _ := time.Parse(time.RFC3339Nano, "2026-01-01T00:00:00.123456789Z")
		page, err = c.Orders.List(ctx, &affinity.ListOrdersRequest{UpdatedAfter: &nanos, Limit: affinity.Int(5)})
		if err == nil {
			note = fmt.Sprintf("%d", len(page.Results))
		}
		record("orders.list updatedAfter (nanoseconds)", err, note)

		n := 0
		page, err = c.Orders.List(ctx, &affinity.ListOrdersRequest{Limit: affinity.Int(1), PracticeID: affinity.String(practice.ID)})
		if err == nil {
			it := page.Iterator()
			for it.Next(ctx) {
				n++
			}
			err = it.Err()
		}
		if err == nil && n != 1 {
			err = fmt.Errorf("%d orders for a practice with 1", n)
		}
		record("orders.list (every page, 1 per page)", err, "")

		cancelled, err := c.Orders.Cancel(ctx, &affinity.CancelOrdersRequest{OrderID: order.ID, IdempotencyKey: key()})
		note = ""
		if err == nil {
			note = string(cancelled.Status)
			if cancelled.CancelRequestedAt != nil {
				note += ", waiting for the pharmacy"
			}
		}
		record("orders.cancel", err, note)
	}
	_, err = c.Orders.Get(ctx, &affinity.GetOrdersRequest{OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc"})
	refused[*affinity.NotFoundError]("orders.get for an unknown order raises NotFoundError", err, nil)

	// Webhook endpoints.
	if webhookURL == "" {
		results = append(results, result{"webhooks", false, "AFFINITY_SMOKE_WEBHOOK_URL is not set"})
		return nil
	}
	endpoint, err := c.Webhooks.Create(ctx, &affinity.CreateWebhooksRequest{IdempotencyKey: key(), URL: webhookURL + "/" + key()})
	record("webhooks.create", err, func() string {
		if endpoint == nil {
			return ""
		}
		return endpoint.ID
	}())
	if err != nil {
		return nil
	}
	*endpointIDs = append(*endpointIDs, endpoint.ID)
	{
		page, err := c.Webhooks.List(ctx, &affinity.ListWebhooksRequest{})
		found, n := false, 0
		if err == nil {
			it := page.Iterator()
			for it.Next(ctx) {
				n++
				found = found || it.Current().ID == endpoint.ID
			}
			err = it.Err()
		}
		if err == nil && !found {
			err = errors.New("the new endpoint is not listed")
		}
		record("webhooks.list (every page)", err, fmt.Sprintf("%d", n))
	}
	updatedEndpoint, err := c.Webhooks.Update(ctx, &affinity.UpdateWebhooksRequest{
		WebhookEndpointID: endpoint.ID,
		IdempotencyKey:    key(),
		URL:               webhookURL + "/" + key(),
	})
	record("webhooks.update", err, idOf(updatedEndpoint))
	resent, err := c.Webhooks.ResendFailed(ctx, &affinity.ResendFailedWebhooksRequest{WebhookEndpointID: endpoint.ID, IdempotencyKey: key()})
	note := ""
	if err == nil {
		note = fmt.Sprintf("queued %d", resent.Queued)
	}
	record("webhooks.resendFailed", err, note)
	removed, err := c.Webhooks.Remove(ctx, &affinity.RemoveWebhooksRequest{WebhookEndpointID: endpoint.ID})
	if err == nil && !removed.Deleted {
		err = errors.New("deleted is false")
	}
	if err == nil {
		*endpointIDs = (*endpointIDs)[:0]
	}
	record("webhooks.remove", err, "")

	raw, err := c.Account.WithRawResponse.Get(ctx)
	note = ""
	if err == nil {
		note = raw.Header.Get("X-Request-Id")
		if note == "" {
			err = errors.New("no X-Request-Id header")
		}
	}
	record("WithRawResponse gives the request ID", err, note)
	return nil
}

// idOf returns the ID of a record with an ID field, or "" for nil.
func idOf[T interface{ GetID() string }](record T) string {
	return record.GetID()
}
