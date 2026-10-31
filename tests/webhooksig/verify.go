// Package webhooksig holds the webhook signature check shown in README.md, so the example is
// compiled and tested. Copy it into your application; it uses only the standard library.
package webhooksig

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
