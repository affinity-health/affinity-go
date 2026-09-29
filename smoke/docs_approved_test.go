package smoke_test
import("context";"errors";"log";affinity "github.com/affinity-health/affinity-go")
func syncPatient(_ any)error{return nil}
func example0()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
patients, err := api.Patients.List(ctx, affinity.PatientListParams{Limit: 20})
if err != nil { return err }

patient, err := api.Patients.Get(ctx, patientID)
if err != nil { return err }

items, err := api.Catalog.Items.List(ctx, affinity.CatalogItemListParams{Limit: 20})
if err != nil { return err }

_ = patients;_ = patient;_ = items
return nil}
func example1()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
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

_ = patients
return nil}
func example2()error{api:=affinity.NewClient("test");ctx:=context.Background();practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
practice := api.ForPractice(practiceID)

patients, err := practice.Patients.List(ctx, affinity.PatientListParams{Limit: 20})
if err != nil { return err }

items, err := practice.Catalog.Items.List(ctx, affinity.CatalogItemListParams{Limit: 20})
if err != nil { return err }

_ = practice;_ = patients;_ = items
return nil}
func example3()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
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

_ = patient;_ = saved
return nil}
func example4()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
_, err = practice.Patients.Delete(ctx, patientID)
if err != nil { return err }


return nil}
func example5()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
order, err := api.Orders.Create(ctx,
    affinity.OrderCreateParams{PatientID: patientID, Prescriptions: draft.Prescriptions},
    affinity.RequestOptions{PracticeID: practiceID, IdempotencyKey: job.CreateOrderKey},
)
if err != nil { return err }

_ = order
return nil}
func example6()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
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

_ = submission
return nil}
func example7()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
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

_ = page;_ = iterator;_ = err
return nil}
func example8()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
_, err = practice.Patients.Get(ctx, patientID)
if err != nil {
    var apiError *affinity.Error
    if errors.As(err, &apiError) {
        log.Printf("status=%d code=%s request=%s retryable=%t retryAfter=%v",
            apiError.Status, apiError.Code, apiError.RequestID,
            apiError.Retryable, apiError.RetryAfter)
    }
    return err
}


return nil}
func example9()error{api:=affinity.NewClient("test");ctx:=context.Background();practice:=api.ForPractice("prac_a");practiceID,patientID,orderID:="prac_a","pat_a","ord_a";draft:=struct{Prescriptions []affinity.OrderCreateParamsPrescriptionsItem}{};job:=struct{CreateOrderKey,SignOrderKey,SubmitOrderKey string}{};review:=struct{PrescriberID,OrderRevision string;SignatureAttestation bool}{};_ = api;_ = ctx;_ = practice;_ = practiceID;_ = patientID;_ = orderID;_ = draft;_ = job;_ = review;var err error;_ = err;
practices, err := api.Practices.List(ctx, affinity.PracticeListParams{Limit: 20})
if err != nil { return err }

selected, err := api.Practices.Get(ctx, practiceID)
if err != nil { return err }

endpoints, err := api.Webhooks.Endpoints.List(
    ctx,
    affinity.WebhookEndpointListParams{Limit: 20},
)
if err != nil { return err }

_ = practices;_ = selected;_ = endpoints
return nil}
