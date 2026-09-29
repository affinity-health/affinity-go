# Go SDK guide

Go applications. Put request statements inside a function that returns an error. [Source repository](https://github.com/affinity-health/affinity-go) · [All SDKs](https://docs.joinaffinityai.com/guides/reference/sdks/)

## Install

```sh
go get github.com/affinity-health/affinity-go@main
```

For reproducible builds, pin the Git dependency to a commit.

## Connect

Set `AFFINITY_API_KEY` to a Test API key on your server. The key selects Test or Live mode. Keep it out of browser and mobile code.

```go
import (
    "context"
    "errors"
    "log"
    "os"

    affinity "github.com/affinity-health/affinity-go"
)

api := affinity.NewClient(os.Getenv("AFFINITY_API_KEY"))
ctx := context.Background()
```

## With a practice key

The key identifies the practice. No practice ID or scoped client is needed.
The resource IDs below come from records in that practice.
Each section is a separate usage example, not one script to concatenate.

```go
patients, err := api.Patients.List(ctx, affinity.PatientListParams{Limit: 20})
if err != nil { return err }

patient, err := api.Patients.Get(ctx, patientID)
if err != nil { return err }

items, err := api.Catalog.Items.List(ctx, affinity.CatalogItemListParams{Limit: 20})
if err != nil { return err }
```

## With a platform key

Pass the target practice with each practice-scoped request. Keep record data separate from request context and idempotency options.

```go
patients, err := api.Patients.List(ctx,
    affinity.PatientListParams{Limit: 20},
    affinity.RequestOptions{PracticeID: practiceID},
)
if err != nil { return err }

_, err = api.Patients.Update(ctx, patientID,
    affinity.PatientUpdateParams{Email: affinity.String("alex@example.com")},
    affinity.RequestOptions{
        PracticeID: practiceID,
    },
)
if err != nil { return err }
```

## Scope a workflow once

A scoped client remembers the practice for subsequent requests. It is immutable; the original client and other scoped clients stay independent.
A conflicting practice ID produces an error. Scoping never grants access to another practice.

```go
practice := api.ForPractice(practiceID)

patients, err := practice.Patients.List(ctx, affinity.PatientListParams{Limit: 20})
if err != nil { return err }

items, err := practice.Catalog.Items.List(ctx, affinity.CatalogItemListParams{Limit: 20})
if err != nil { return err }
```

The following examples use this scoped client. A practice-key client supports the same calls without the scoping step.

## Create, get, and update a patient

Use synthetic Test data. Routine writes generate a fresh idempotency key per call and preserve it during internal retries.
Supply your own persisted key when retrying across calls or process restarts.

```go
patient, err := practice.Patients.Create(ctx, affinity.PatientCreateParams{
    Name: affinity.PatientName{First: "Alex", Last: "Example"},
    DateOfBirth: "1990-01-01",
})
if err != nil { return err }

saved, err := practice.Patients.Get(ctx, patient.ID)
if err != nil { return err }

_, err = practice.Patients.Update(ctx, patient.ID, affinity.PatientUpdateParams{
    Email: affinity.String("alex@example.com"),
})
if err != nil { return err }

_, err = practice.Patients.Update(ctx, patient.ID, affinity.PatientUpdateParams{
    Status: affinity.String("archived"),
})
if err != nil { return err }
```

The SDK maps `archived` to the API’s `inactive` status. Returned records use `inactive`.

Archive patients whose records you need to retain. Permanent deletion is available only for patients without order history. No explicit idempotency key is needed.

```go
_, err = practice.Patients.Delete(ctx, patientID)
if err != nil { return err }
```

## Create an order draft

`draft` is your application's prepared prescription data, using catalog and prescribing options from this practice.
An order contains 1–20 complete prescriptions for one patient. This example creates an unsigned draft.
It shows a platform call without a scoped client: practice context and the persisted key belong together in request options.

`job` is your persisted workflow record. Generate and save a unique key for each action before making its first request.

```go
order, err := api.Orders.Create(ctx,
    affinity.OrderCreateParams{PatientID: patientID, Prescriptions: draft.Prescriptions},
    affinity.RequestOptions{PracticeID: practiceID, IdempotencyKey: job.CreateOrderKey},
)
if err != nil { return err }
```

## Sign and submit

`review` is your saved clinician review and signing consent for this exact order.
Store the reviewed revision, authorized prescriber ID, and explicit attestation together.
Your API key needs `orders:sign`. Never infer consent or automatically replace a stale revision.

```go
_, err = practice.Orders.Sign(ctx, orderID,
    affinity.OrderSignParams{
        Prescriber: affinity.PrescriberSelector{ID: review.PrescriberID},
        ExpectedRevision: review.OrderRevision,
        SignatureAttestation: review.SignatureAttestation,
    },
    affinity.RequestOptions{IdempotencyKey: job.SignOrderKey},
)
if err != nil { return err }

submission, err := practice.Orders.Submit(ctx, orderID,
    affinity.RequestOptions{IdempotencyKey: job.SubmitOrderKey},
)
if err != nil { return err }
```

Use separate keys for creating, signing, and submitting. After an uncertain response, retry the same action with the same key and unchanged data.
A revision conflict requires renewed clinician review before another signing attempt.

Submission means queued, not accepted by the pharmacy. Inspect the result and track order events or webhooks.
After a reported partial submission failure, retry only the unconfirmed send with a new submission key.

## Read more than one page

The list method returns one page. Pass the last record's ID to request the next page.
The iterator fetches pages as you consume records; it does not load the full collection into memory.
`syncPatient` or its language equivalent represents your application's record handler.

```go
page, err := practice.Patients.List(ctx, affinity.PatientListParams{Limit: 20})
if err != nil { return err }

if page.HasMore && len(page.Data) > 0 {
    next, err := practice.Patients.List(ctx, affinity.PatientListParams{
        Limit: 20, StartingAfter: page.Data[len(page.Data)-1].ID,
    })
    if err != nil { return err }
    _ = next
}

iterator := practice.Patients.Iterate(ctx, affinity.PatientListParams{Limit: 100})
for iterator.Next() {
    if err := syncPatient(iterator.Value()); err != nil { return err }
}
if err := iterator.Err(); err != nil { return err }
```

## Handle errors

API failures expose status, code, request ID, retryability, and an optional retry delay in seconds.
Log those fields without logging patient data or credentials. Transport failures remain distinguishable from API responses.

```go
_, err := practice.Patients.Get(ctx, patientID)
if err != nil {
    var apiError *affinity.Error
    if errors.As(err, &apiError) {
        log.Printf("status=%d code=%s request=%s retryable=%t retryAfter=%v",
            apiError.Status, apiError.Code, apiError.RequestID,
            apiError.Retryable, apiError.RetryAfter)
    }
    return err
}
```

Retryability is a transport hint, not permission to repeat a clinical action with a new key.
Keep the same key and body for an uncertain write. Validation and authorization errors require a corrected request.
See [API errors](https://docs.joinaffinityai.com/errors/) for recovery guidance.

## Platform directory and webhooks

Use the root platform client to list its practices and webhook endpoints. These calls do not need a target practice or an idempotency key.
The webhook list belongs to the platform itself. Access to another organization's endpoints still requires an explicit grant.

```go
practices, err := api.Practices.List(ctx, affinity.PracticeListParams{Limit: 20})
if err != nil { return err }

selected, err := api.Practices.Get(ctx, practiceID)
if err != nil { return err }

endpoints, err := api.Webhooks.Endpoints.List(
    ctx,
    affinity.WebhookEndpointListParams{Limit: 20},
)
if err != nil { return err }
```

## More resources

Use the same conventions for addresses, allergies, locations, team members, and nested order resources.
[API reference](https://docs.joinaffinityai.com/api/) · [Webhooks](https://docs.joinaffinityai.com/guides/webhooks/)
