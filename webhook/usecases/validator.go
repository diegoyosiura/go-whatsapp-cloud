package usecases

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

// ValidateSignature verifies if the provided payload matches the X-Hub-Signature-256
// hash sent by Meta/WhatsApp, using the webhook secret.
func ValidateSignature(payload []byte, signatureHeader string, secret string) error {
	if signatureHeader == "" || !strings.HasPrefix(signatureHeader, "sha256=") {
		return errors.New("invalid signature format")
	}

	providedMAC := strings.TrimPrefix(signatureHeader, "sha256=")

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(providedMAC), []byte(expectedMAC)) {
		return errors.New("invalid signature")
	}

	return nil
}
