package usecases

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"testing"
)

func TestValidateSignature(t *testing.T) {
	secret := "my_test_secret"
	payload := []byte(`{"object":"whatsapp_business_account","entry":[{"id":"123"}]}`)

	// Generate a valid signature for the test
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	validSignature := "sha256=" + expectedMAC

	tests := []struct {
		name        string
		payload     []byte
		signature   string
		secret      string
		expectedErr error
	}{
		{
			name:        "Valid Signature",
			payload:     payload,
			signature:   validSignature,
			secret:      secret,
			expectedErr: nil,
		},
		{
			name:        "Invalid Signature Content",
			payload:     payload,
			signature:   "sha256=invalidhash",
			secret:      secret,
			expectedErr: errors.New("invalid signature"),
		},
		{
			name:        "Missing prefix",
			payload:     payload,
			signature:   expectedMAC, // missing "sha256="
			secret:      secret,
			expectedErr: errors.New("invalid signature format"),
		},
		{
			name:        "Empty Signature",
			payload:     payload,
			signature:   "",
			secret:      secret,
			expectedErr: errors.New("invalid signature format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSignature(tt.payload, tt.signature, tt.secret)

			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedErr)
				} else if err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}
