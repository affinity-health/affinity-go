// Tests client.New's defaults against a local server: retries reuse the idempotency key, options
// replace the defaults, and the timeout ends a slow attempt.
package defaults

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	affinity "github.com/affinity-health/affinity-go"
	"github.com/affinity-health/affinity-go/client"
	"github.com/affinity-health/affinity-go/option"
)

type recorder struct {
	mu   sync.Mutex
	keys []string
}

// server answers 503 `failures` times, then a practice.
func server(t *testing.T, failures int, delay time.Duration) (*httptest.Server, *recorder) {
	t.Helper()
	rec := &recorder{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.mu.Lock()
		rec.keys = append(rec.keys, r.Header.Get("Idempotency-Key"))
		n := len(rec.keys)
		rec.mu.Unlock()
		time.Sleep(delay)
		w.Header().Set("Content-Type", "application/json")
		if n <= failures {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"type":"t","title":"t","status":503,"code":"service_unavailable","detail":"d","instance":"/v1/practices","requestId":"r"}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":"prac_01j2y8m6jcc9tt24af5pw9x1bc","name":"Synthetic","phone":"+15125550100","liveEnabled":false,"createdAt":"2026-10-31T00:00:00Z","updatedAt":"2026-10-31T00:00:00Z","mode":"test","type":"practice"}`))
	}))
	t.Cleanup(srv.Close)
	return srv, rec
}

func create(c *client.Client) error {
	_, err := c.Practices.Create(context.Background(), &affinity.CreatePracticesRequest{
		IdempotencyKey: "saved-key",
		Name:           "Synthetic",
		Phone:          "+15125550100",
		Address:        &affinity.PracticeAddressInput{Line1: "100 Congress Ave", City: "Austin", State: "TX", PostalCode: "78701"},
	})
	return err
}

func TestNewRetriesTwiceWithTheSameKey(t *testing.T) {
	srv, rec := server(t, 2, 0)
	if err := create(client.New("sk_test_x", option.WithBaseURL(srv.URL))); err != nil {
		t.Fatalf("create: %v", err)
	}
	if len(rec.keys) != 3 {
		t.Fatalf("attempts = %d, want 3", len(rec.keys))
	}
	for _, key := range rec.keys {
		if key != "saved-key" {
			t.Fatalf("keys = %v, want saved-key on every attempt", rec.keys)
		}
	}
}

func TestNewGivesUpAfterTwoRetries(t *testing.T) {
	srv, rec := server(t, 5, 0)
	if err := create(client.New("sk_test_x", option.WithBaseURL(srv.URL))); err == nil {
		t.Fatal("want an error after three 503s")
	}
	if len(rec.keys) != 3 {
		t.Fatalf("attempts = %d, want 3", len(rec.keys))
	}
}

func TestOptionsReplaceTheDefaults(t *testing.T) {
	srv, rec := server(t, 5, 0)
	_ = create(client.New("sk_test_x", option.WithBaseURL(srv.URL), option.WithoutRetries()))
	if len(rec.keys) != 1 {
		t.Fatalf("attempts = %d, want 1 with WithoutRetries", len(rec.keys))
	}
}

func TestTimeoutEndsASlowAttempt(t *testing.T) {
	srv, _ := server(t, 0, 2*time.Second)
	started := time.Now()
	err := create(client.New("sk_test_x", option.WithBaseURL(srv.URL), option.WithoutRetries(),
		option.WithHTTPClient(&http.Client{Timeout: 200 * time.Millisecond})))
	if err == nil || time.Since(started) > time.Second {
		t.Fatalf("err = %v after %v, want a timeout within 1s", err, time.Since(started))
	}
}

func TestNewUsesTheDefaultTimeout(t *testing.T) {
	if client.DefaultTimeout != 60*time.Second || client.DefaultMaxAttempts != 3 {
		t.Fatalf("defaults = %v, %d", client.DefaultTimeout, client.DefaultMaxAttempts)
	}
}
