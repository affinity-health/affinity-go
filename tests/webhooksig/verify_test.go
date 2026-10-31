package webhooksig

import (
	"testing"
	"time"
)

// The expected signature was computed independently of this code:
//
//	printf '%s' '1793491200.{"id":"evt_01j2y8m6jcc9tt24af5pw9x1bc","type":"order.updated"}' \
//	  | openssl dgst -sha256 -hmac 'whsec_smoke_example_secret'
const (
	secret    = "whsec_smoke_example_secret"
	body      = `{"id":"evt_01j2y8m6jcc9tt24af5pw9x1bc","type":"order.updated"}`
	signature = "aeec8fea545cbb30eae06f24ab18d0ef5f570a8517b3b696f257b00b7e2ffd46"
)

var sentAt = time.Unix(1793491200, 0)

func TestVerifyAffinitySignature(t *testing.T) {
	header := "t=1793491200,v1=" + signature
	cases := []struct {
		name   string
		header string
		body   string
		secret string
		now    time.Time
		ok     bool
	}{
		{"valid", header, body, secret, sentAt.Add(10 * time.Second), true},
		{"valid among other v1 values", "t=1793491200,v1=00ff,v1=" + signature, body, secret, sentAt, true},
		{"tampered body", header, body + " ", secret, sentAt, false},
		{"wrong secret", header, body, "whsec_other", sentAt, false},
		{"older than 5 minutes", header, body, secret, sentAt.Add(301 * time.Second), false},
		{"more than 5 minutes ahead", header, body, secret, sentAt.Add(-301 * time.Second), false},
		{"uppercase hex", "t=1793491200,v1=AEEC8FEA545CBB30EAE06F24AB18D0EF5F570A8517B3B696F257B00B7E2FFD46", body, secret, sentAt, false},
		{"no timestamp", "v1=" + signature, body, secret, sentAt, false},
		{"no signature", "t=1793491200", body, secret, sentAt, false},
	}
	for _, c := range cases {
		err := VerifyAffinitySignature(c.header, []byte(c.body), c.secret, c.now)
		if (err == nil) != c.ok {
			t.Errorf("%s: got error %v, want ok=%v", c.name, err, c.ok)
		}
	}
}
