// A response with a value or field this SDK does not know decodes without an error. The contract
// keeps response objects and value lists open: a later API version can add a status, a category or
// a field. The fixture has an unknown order, prescription and history status, an unknown allergy
// category and an added field at the root and in nested objects.
package openvalues

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	affinity "github.com/affinity-health/affinity-go"
	"github.com/affinity-health/affinity-go/client"
	"github.com/affinity-health/affinity-go/option"
)

func fixture(t *testing.T) string {
	t.Helper()
	body, err := os.ReadFile("testdata/order-unknown-values.json")
	if err != nil {
		t.Fatal(err)
	}
	return string(body)
}

func serve(t *testing.T, bodies ...string) *client.Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(bodies) == 0 {
			t.Errorf("unexpected request %s %s", r.Method, r.URL)
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(bodies[0]))
		bodies = bodies[1:]
	}))
	t.Cleanup(srv.Close)
	return client.New("sk_test_synthetic", option.WithBaseURL(srv.URL), option.WithoutRetries())
}

func TestOrdersGetKeepsUnknownValuesAndAddedFields(t *testing.T) {
	order, err := serve(t, fixture(t)).Orders.Get(context.Background(),
		&affinity.GetOrdersRequest{OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc"})
	if err != nil {
		t.Fatal(err)
	}
	if order.Status != "on_hold" || order.Prescriptions[0].Status != "on_hold" {
		t.Errorf("status %q, prescription status %q", order.Status, order.Prescriptions[0].Status)
	}
	if order.StatusHistory[0].Status != affinity.StatusHistoryEntryStatusSubmitted ||
		order.StatusHistory[1].Status != "on_hold" {
		t.Errorf("history %q %q", order.StatusHistory[0].Status, order.StatusHistory[1].Status)
	}
	allergies := order.Patient.Allergies
	if allergies[0].Category != affinity.AllergyCategoryDrug || allergies[1].Category != "dust" {
		t.Errorf("categories %q %q", allergies[0].Category, allergies[1].Category)
	}
	if added, ok := order.GetExtraProperties()["addedField"].(map[string]interface{}); !ok || added["nested"] != true {
		t.Errorf("order addedField %#v", order.GetExtraProperties())
	}
	if order.Patient.GetExtraProperties()["addedField"] != "patient" {
		t.Errorf("patient addedField %#v", order.Patient.GetExtraProperties())
	}
	if order.Prescriptions[0].Medication.GetExtraProperties()["addedField"] != "medication" {
		t.Errorf("medication addedField %#v", order.Prescriptions[0].Medication.GetExtraProperties())
	}
}

func TestOrdersListIteratesUnknownValues(t *testing.T) {
	page := func(cursor string) string {
		next := "null"
		if cursor != "" {
			next = `"` + cursor + `"`
		}
		return `{"object":"list","url":"/v1/orders","data":[` + fixture(t) + `],"hasMore":` +
			map[bool]string{true: "true", false: "false"}[cursor != ""] + `,"nextCursor":` + next +
			`,"addedField":"list"}`
	}
	pages, err := serve(t, page("cursor-a"), page("")).Orders.List(context.Background(), &affinity.ListOrdersRequest{})
	if err != nil {
		t.Fatal(err)
	}
	var statuses []string
	iter := pages.Iterator()
	for iter.Next(context.Background()) {
		statuses = append(statuses, string(iter.Current().Status))
	}
	if err := iter.Err(); err != nil {
		t.Fatal(err)
	}
	if strings.Join(statuses, ",") != "on_hold,on_hold" {
		t.Errorf("statuses %v", statuses)
	}
}

// The known values stay typed: each has a named constant of the enum type.
var _ = []affinity.OrderStatus{affinity.OrderStatusAwaitingApproval, affinity.OrderStatusFailed}
var _ = []affinity.AllergyCategory{affinity.AllergyCategoryDrug, affinity.AllergyCategoryPollen}
