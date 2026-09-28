package smoke_test

import (
    "context"
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    affinity "github.com/affinity-health/affinity-go"
    "github.com/affinity-health/affinity-go/client"
    "github.com/affinity-health/affinity-go/core"
    "github.com/affinity-health/affinity-go/option"
)

func ptr[T any](value T) *T { return &value }

func TestHeadersQueryAndResponse(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/v1/orders" || r.Method != "GET" { t.Errorf("unexpected request: %s %s", r.Method, r.URL) }
        for name, value := range map[string]string{
            "x-affinity-api-key": "synthetic-key", "Affinity-Version": "2026-09-28",
            "Affinity-Actor-Id": "user-synthetic", "Affinity-Actor-Type": "user",
        } { if r.Header.Get(name) != value { t.Errorf("incorrect %s", name) } }
        if r.URL.Query().Get("startingAfter") != "ord_cursor" || r.URL.Query().Get("limit") != "2" {
            t.Errorf("incorrect query: %s", r.URL.RawQuery)
        }
        if r.URL.Query().Has("patientId") { t.Error("omitted patientId was sent") }
        w.Header().Set("Content-Type", "application/json")
        io.WriteString(w, `{"object":"list","data":[],"hasMore":false,"url":"/v1/orders"}`)
    }))
    defer server.Close()
    sdk := client.NewClient(option.WithAPIKey("synthetic-key"), option.WithAffinityVersion(ptr("2026-09-28")),
        option.WithBaseURL(server.URL), option.WithoutRetries())
    page, err := sdk.Orders.ListOrders(context.Background(), &affinity.ListOrdersRequest{
        StartingAfter: ptr("ord_cursor"), Limit: ptr(2), AffinityActorID: ptr("user-synthetic"), AffinityActorType: ptr("user"),
    })
    if err != nil { t.Fatal(err) }
    if page.HasMore || len(page.Data) != 0 { t.Fatalf("incorrect response: %+v", page) }
}

func TestKeyedRetryAndError(t *testing.T) {
    var bodies []string
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/v1/orders" || r.Method != "POST" { t.Errorf("unexpected request: %s %s", r.Method, r.URL) }
        if r.Header.Get("Idempotency-Key") != "stable-synthetic-key" { t.Error("missing stable idempotency key") }
        body, _ := io.ReadAll(r.Body)
        bodies = append(bodies, string(body))
        var payload map[string]any
        if err := json.Unmarshal(body, &payload); err != nil { t.Error(err) }
        if payload["practiceId"] != "prac_synthetic" || payload["patientId"] != "pat_synthetic" {
            t.Errorf("incorrect body: %s", body)
        }
        w.Header().Set("Content-Type", "application/problem+json")
        if len(bodies) == 1 { w.WriteHeader(503) } else { w.WriteHeader(422) }
        io.WriteString(w, `{"type":"about:blank","title":"Invalid request","status":422,"detail":"Synthetic validation failure"}`)
    }))
    defer server.Close()
    sdk := client.NewClient(option.WithAPIKey("synthetic-key"), option.WithAffinityVersion(ptr("2026-09-28")),
        option.WithBaseURL(server.URL), option.WithMaxAttempts(2))
    _, err := sdk.Orders.CreateOrder(context.Background(), &affinity.CreateOrderRequest{
        IdempotencyKey: "stable-synthetic-key", PracticeID: "prac_synthetic", PatientID: ptr("pat_synthetic"),
        Prescriptions: []*affinity.CreateOrderRequestPrescriptionsItem{},
    })
    var apiError *core.APIError
    if !errors.As(err, &apiError) || apiError.StatusCode != 422 || !strings.Contains(err.Error(), "Synthetic validation failure") {
        t.Fatalf("incorrect error: %v", err)
    }
    if len(bodies) != 2 || bodies[0] != bodies[1] { t.Fatalf("retry changed request: %v", bodies) }
}

func TestCanceledContext(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    sdk := client.NewClient(option.WithAPIKey("synthetic-key"), option.WithBaseURL("http://127.0.0.1:1"), option.WithoutRetries())
    _, err := sdk.APIKeys.GetAPIAccess(ctx)
    if !errors.Is(err, context.Canceled) { t.Fatalf("expected cancellation: %v", err) }
}
