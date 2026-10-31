# Affinity Go SDK

Typed server-side Go client for the Affinity API. Requires Go 1.21 or later. It targets API version `2026-10-28` and sends it in `Affinity-Version` on every request.

Keep API keys on the server. Never put them in browser code, mobile applications or public configuration.

## Install

Install from GitHub at tag `v0.4.1`. The module is not published to a registry other than GitHub.

```sh
go get github.com/affinity-health/affinity-go@v0.4.1
```

## Use

Create the client with `client.New(key, options...)`. Use a Test key (`sk_test_…`) while you build: Test orders go to a test pharmacy.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	affinity "github.com/affinity-health/affinity-go"
	"github.com/affinity-health/affinity-go/client"
	"github.com/affinity-health/affinity-go/option"
)

func main() {
	ctx := context.Background()
	c := client.New(os.Getenv("AFFINITY_API_KEY"))

	account, err := c.Account.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if account.Mode != affinity.AccountModeTest {
		log.Fatal("use a Test key for this example")
	}

	page, err := c.Medications.List(ctx, &affinity.ListMedicationsRequest{Limit: affinity.Int(25)})
	if err != nil {
		log.Fatal(err)
	}
	it := page.Iterator()
	for it.Next(ctx) {
		fmt.Println(it.Current().Name, it.Current().Pharmacy.Name)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

The key sets the organization and the Test or Live mode. Each resource is a field of the client: `Account`, `Feedback`, `Medications`, `Orders`, `Practices` and `Webhooks`. Every method takes a `context.Context` and one request struct. `affinity.String`, `affinity.Int` and the other helpers in `pointer.go` make pointers for optional fields.

### Idempotency keys

Every write requires `IdempotencyKey` in its request struct (`OrderInput`, `CreatePracticesRequest`, `CancelOrdersRequest` and the others). The SDK does not make one for you. Create the key when the order is reviewed, save it with the order, then send the request. To retry, send the same key and the same body: the API returns the first result instead of creating a second order. The client's automatic retries send the same request, with the same key.

```go
order, err := c.Orders.Create(ctx, &affinity.OrderInput{
	IdempotencyKey: savedKey, // saved with the reviewed order
	PracticeID:     affinity.String("prac_..."),
	// Patient, Prescriber, Prescriptions, Shipping, SignatureAttestation: true
})
```

### Pagination

List methods return a `*core.Page`. `page.Results` is one page. `page.Iterator()` reads every item and fetches each next page as needed; check `it.Err()` after the loop. `page.GetNextPage(ctx)` returns `core.ErrNoPages` after the last page. Keep filters the same while you page.

`ListOrdersRequest.UpdatedAfter` is a `*time.Time`. The SDK sends it in RFC 3339 with milliseconds, so finer precision is cut to milliseconds.

### Errors

A non-2xx response returns an error of the type for its status: `*affinity.BadRequestError` (400), `*affinity.UnauthorizedError` (401), `*affinity.ForbiddenError` (403), `*affinity.NotFoundError` (404), `*affinity.ConflictError` (409), `*affinity.UnprocessableEntityError` (422), `*affinity.TooManyRequestsError` (429), `*affinity.InternalServerError` (500), `*affinity.ServiceUnavailableError` (503) and others. Each embeds `*core.APIError`, which has `StatusCode` and `Header`. `GetBody()` returns the problem body; a 422 lists each refused field in `Issues`.

```go
var invalid *affinity.UnprocessableEntityError
if errors.As(err, &invalid) {
	for _, issue := range invalid.GetBody().GetIssues() {
		fmt.Println(issue.Field, issue.Message) // for example prescriptions[0].quantity
	}
}
```

Every problem body has `RequestID`. For a successful call, `WithRawResponse` gives the headers: `raw, err := c.Account.WithRawResponse.Get(ctx)`, then `raw.Header.Get("X-Request-Id")` and `raw.Body`.

### Defaults

- Base URL `https://api.affinityrx.com` (`affinity.Environments.Default`). Change it with `option.WithBaseURL`.
- `Affinity-Version: 2026-10-28` on every request.
- Retries: 2 retries (3 attempts) for 408, 429 and 5xx responses, with exponential backoff that honors `Retry-After`. A retry sends the same request, with the same `IdempotencyKey`. Network errors are not retried: retry them yourself with the same key. Change it with `option.WithMaxAttempts(n)` or turn it off with `option.WithoutRetries()`.
- Timeout: 60 seconds per attempt (`client.DefaultTimeout`). Replace it with `option.WithHTTPClient(&http.Client{Timeout: ...})`, or set a deadline on the context.

`client.New` sets these defaults; options you pass after the key replace them. Fern's `client.NewClient` is also available, but it has no timeout and makes 1 retry.

Options passed to `client.New` apply to every request; options passed to a method apply to that call only.

See [docs/reference.md](docs/reference.md) for every method and https://docs.affinityrx.com for the API.

## Verify webhooks

Affinity signs each webhook delivery with the endpoint secret (`whsec_…`) that `Webhooks.Create` returns. The `Affinity-Signature` header is `t=<unix seconds>,v1=<hex>`. The signature is the lowercase hex HMAC-SHA256 of `<t>.` followed by the raw request body. Fern does not generate a verifier; this standard-library function does the check. Pass the body bytes exactly as received, before any JSON decoding.

```go
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"
)

// VerifyAffinitySignature checks an Affinity-Signature header (t=<unix>,v1=<hex>) against the raw
// request body and the endpoint secret (whsec_…). It returns nil when one v1 signature matches and
// the timestamp is within 5 minutes of now.
func VerifyAffinitySignature(header string, body []byte, secret string, now time.Time) error {
	var timestamp string
	var signatures []string
	for _, part := range strings.Split(header, ",") {
		key, value, ok := strings.Cut(strings.TrimSpace(part), "=")
		if !ok {
			continue
		}
		switch key {
		case "t":
			timestamp = value
		case "v1":
			signatures = append(signatures, value)
		}
	}
	seconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil || len(signatures) == 0 {
		return errors.New("affinity: malformed Affinity-Signature header")
	}
	if age := now.Sub(time.Unix(seconds, 0)); age > 5*time.Minute || age < -5*time.Minute {
		return errors.New("affinity: signature timestamp is outside the 5-minute tolerance")
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "."))
	mac.Write(body)
	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	for _, signature := range signatures {
		if hmac.Equal(expected, []byte(signature)) {
			return nil
		}
	}
	return errors.New("affinity: no matching signature")
}
```

In an HTTP handler, read the body with `io.ReadAll(r.Body)`, call `VerifyAffinitySignature(r.Header.Get("Affinity-Signature"), body, secret, time.Now())`, and reply 400 when it returns an error. The same function is in [tests/webhooksig/verify.go](tests/webhooksig/verify.go), tested against a signature computed with `openssl dgst -sha256 -hmac` and against a changed body.

## Build and verify

[Fern](https://buildwithfern.com) generates the Go code locally, in Docker, from `spec/affinity.openapi.json`, the committed Affinity API contract. No Fern account is used. [generation.json](generation.json) records the contract version and hash, the Fern version and the generator version.

```sh
# Build, vet and run the unit tests that need no credentials or services.
docker run --rm -v "$PWD":/src -w /src golang:1.23 \
  sh -c 'go build ./... && go vet ./... && go test $(go list ./... | grep -v "_test$")'

# Regenerate after the contract changes (needs Docker, Node and jq).
scripts/generate.sh

# Regenerate and fail if any generated file differs from the committed one.
scripts/generate-check.sh

# Every operation against a real API, Test keys only. The env file sets AFFINITY_API_KEY,
# AFFINITY_BASE_URL and AFFINITY_SMOKE_WEBHOOK_URL.
docker run --rm --env-file smoke.env -v "$PWD":/src -w /src golang:1.23 go run ./tests/smoke
```

The unit-test command leaves out the `*/…_test` packages: Fern's wire tests there need a WireMock server (`wiremock/docker-compose.test.yml`). The smoke test refuses a key that does not start with `sk_test_`. It creates a Test practice and a Test order, cancels every order it creates and removes every webhook endpoint it creates. The Test practice stays: the API cannot remove a practice.
