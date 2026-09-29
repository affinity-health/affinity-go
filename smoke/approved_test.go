package smoke_test
import("context";"encoding/json";"net/http";"testing";"errors";affinity "github.com/affinity-health/affinity-go")
func TestApprovedInterface(t *testing.T){
 base:="http://127.0.0.1:5199/go-practice-retry";resp,e:=http.Get(base+"/reset");if e!=nil{t.Fatal(e)};resp.Body.Close()
 api:=affinity.NewClient("test",affinity.ClientOptions{BaseURL:base,MaxRetries:1});ctx:=context.Background()
 patient,e:=api.Patients.Create(ctx,affinity.PatientCreateParams{Name:affinity.PatientName{First:"Alex",Last:"Example"},DateOfBirth:"1990-01-01"});if e!=nil{t.Fatal(e)};if patient.ID!="pat_a"{t.Fatal("patient")}
 _,e=api.Patients.Update(ctx,patient.ID,affinity.PatientUpdateParams{NullFields:[]string{"email"}});if e!=nil{t.Fatal(e)}
 _,e=api.Patients.Delete(ctx,patient.ID);if e!=nil{t.Fatal(e)}
 iterator:=api.Patients.Iterate(ctx,affinity.PatientListParams{Limit:1,Query:"Alex"});var ids []string;for iterator.Next(){ids=append(ids,iterator.Value().ID)};if iterator.Err()!=nil{t.Fatal(iterator.Err())};if len(ids)!=2||ids[0]!="pat_a"||ids[1]!="pat_b"{t.Fatal(ids)}
 if _,e=api.Patients.Get(ctx,"pat_a",affinity.RequestOptions{PracticeID:"prac_b"});e==nil{t.Fatal("mismatch accepted")}
 if _,e=api.Orders.Submit(ctx,"ord_a");e==nil{t.Fatal("missing key accepted")}
 _,e=api.Orders.Sign(ctx,"ord_a",affinity.OrderSignParams{Prescriber:affinity.PrescriberSelector{ID:"prov_a"},ExpectedRevision:"rev_reviewed",SignatureAttestation:true},affinity.RequestOptions{IdempotencyKey:"sign_job"});if e!=nil{t.Fatal(e)}
 _,e=api.Orders.Submit(ctx,"ord_a",affinity.RequestOptions{IdempotencyKey:"submit_job"});if e!=nil{t.Fatal(e)}
 _,e=api.Patients.Get(ctx,"pat_error");var apiError *affinity.Error;if !errors.As(e,&apiError)||apiError.Status!=429||apiError.Code!="rate_limited"||apiError.RequestID!="req_a"||apiError.RetryAfter==nil||*apiError.RetryAfter!=0{t.Fatalf("error: %+v",e)}
 resp,e=http.Get(base+"/trace");if e!=nil{t.Fatal(e)};defer resp.Body.Close();var trace []struct{Path string;Method string;Key string;Body map[string]any};if e=json.NewDecoder(resp.Body).Decode(&trace);e!=nil{t.Fatal(e)};access:=0;var writes []string;for _,r:=range trace{if r.Path=="/v1/auth/access"{access++};if r.Method=="PATCH"{writes=append(writes,r.Key);v,ok:=r.Body["email"];if !ok||v!=nil{t.Fatal("null was not preserved")}}};if access!=1||len(writes)!=2||writes[0]!=writes[1]{t.Fatal("cache or retry keys")}
 for _,scoped:=range []*affinity.Client{api.ForPractice(""),api.ForPractice("prac_a").ForPractice("prac_b")} {if _,e=scoped.Patients.Get(ctx,"pat_a");e==nil{t.Fatal("invalid scope accepted")}}
 canceled,cancel:=context.WithCancel(ctx);cancel();if _,e=api.Patients.Get(canceled,"pat_a");!errors.Is(e,context.Canceled){t.Fatal("cancellation")}
}
