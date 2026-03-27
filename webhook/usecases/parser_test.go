package usecases

import (
	"strings"
	"testing"
)

func TestParsePayload(t *testing.T) {
	validJSON := []byte(`{
		"object": "whatsapp_business_account",
		"entry": [
			{
				"id": "WHATSAPP_BUSINESS_ACCOUNT_ID",
				"changes": [
					{
						"value": {
							"messaging_product": "whatsapp",
							"metadata": {
								"display_phone_number": "16505551111",
								"phone_number_id": "123456123"
							},
							"messages": [
								{
									"from": "16505552222",
									"id": "wamid.HBgLMTY1MDU1NTIyMjIVAgASG ...",
									"timestamp": "1603059201",
									"type": "text",
									"text": {
										"body": "Hello this is a test"
									}
								}
							]
						},
						"field": "messages"
					}
				]
			}
		]
	}`)

	invalidJSON := []byte(`{"object": "whatsapp_business_account", "entry": [}`)

	t.Run("Valid Payload", func(t *testing.T) {
		payload, err := ParsePayload(validJSON)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if payload.Object != "whatsapp_business_account" {
			t.Errorf("expected object 'whatsapp_business_account', got '%s'", payload.Object)
		}

		if len(payload.Entry) != 1 {
			t.Fatalf("expected 1 entry, got %d", len(payload.Entry))
		}

		entry := payload.Entry[0]
		if entry.ID != "WHATSAPP_BUSINESS_ACCOUNT_ID" {
			t.Errorf("expected entry ID 'WHATSAPP_BUSINESS_ACCOUNT_ID', got '%s'", entry.ID)
		}

		if len(entry.Changes) != 1 {
			t.Fatalf("expected 1 change, got %d", len(entry.Changes))
		}

		value := entry.Changes[0].Value
		if value.Metadata.PhoneNumberID != "123456123" {
			t.Errorf("expected phone number id '123456123', got '%s'", value.Metadata.PhoneNumberID)
		}

		if len(value.Messages) != 1 {
			t.Fatalf("expected 1 message, got %d", len(value.Messages))
		}

		if value.Messages[0].Text == nil || value.Messages[0].Text.Body != "Hello this is a test" {
			t.Errorf("expected message text body 'Hello this is a test'")
		}
	})

	t.Run("Invalid Payload", func(t *testing.T) {
		_, err := ParsePayload(invalidJSON)
		if err == nil {
			t.Fatalf("expected error for invalid JSON, got none")
		}
		if !strings.Contains(err.Error(), "failed to parse") {
			t.Errorf("expected error to mention parsing failure, got: %v", err)
		}
	})
}
