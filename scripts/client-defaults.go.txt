// Written by hand: scripts/generate.sh copies scripts/client-defaults.go.txt to client/defaults.go after Fern
// generates the package. Fern's Go generator has no setting for these defaults.

package client

import (
	"net/http"
	"time"

	"github.com/affinity-health/affinity-go/option"
)

const (
	// DefaultTimeout limits one attempt of a request made by a client from New.
	DefaultTimeout = 60 * time.Second
	// DefaultMaxAttempts is the first attempt and two retries, for 408, 429 and 5xx responses.
	// A retry sends the same request, with the same Idempotency-Key.
	DefaultMaxAttempts = 3
)

// New returns a client for the API key token with a 60-second timeout per attempt and two
// retries. Options given after the token replace these defaults, for example
// option.WithHTTPClient, option.WithMaxAttempts or option.WithoutRetries.
//
// NewClient, which Fern generates, uses http.DefaultClient (no timeout) and one retry.
func New(token string, opts ...option.RequestOption) *Client {
	defaults := []option.RequestOption{
		option.WithToken(token),
		option.WithHTTPClient(&http.Client{Timeout: DefaultTimeout}),
		option.WithMaxAttempts(DefaultMaxAttempts),
	}
	return NewClient(append(defaults, opts...)...)
}
