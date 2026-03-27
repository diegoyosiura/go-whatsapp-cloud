package services

import (
	"context"
	"testing"
	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/domain"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type mockEventHandler struct {
	messages []domain.Message
	statuses []domain.Status
}

func (m *mockEventHandler) OnMessage(ctx context.Context, msg domain.Message) error {
	m.messages = append(m.messages, msg)
	return nil
}

func (m *mockEventHandler) OnStatus(ctx context.Context, status domain.Status) error {
	m.statuses = append(m.statuses, status)
	return nil
}

func TestWebhookService_ProcessEvent(t *testing.T) {
	secret := "my_test_secret"
	
	// Validação de assinatura será exigida. 
	// A função parse será chamada internamente, então payload tem que ser válido JSON.
	payload := []byte(`{
		"object": "whatsapp_business_account",
		"entry": [
			{
				"id": "WHATSAPP_ID",
				"changes": [
					{
						"value": {
							"messages": [
								{
									"from": "16505551111",
									"id": "wamid.HBLMTY1MDU1",
									"type": "text",
									"text": {"body": "Test message"}
								}
							]
						},
						"field": "messages"
					}
				]
			}
		]
	}`)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	validSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	svc := NewWebhookService(secret)
	handler := &mockEventHandler{}
	svc.RegisterHandler(handler)

	t.Run("Valid routing to OnMessage", func(t *testing.T) {
		err := svc.ProcessEvent(context.Background(), payload, validSignature)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(handler.messages) != 1 {
			t.Fatalf("expected 1 message processed, got %d", len(handler.messages))
		}

		if handler.messages[0].Text.Body != "Test message" {
			t.Errorf("expected text body 'Test message', got %s", handler.messages[0].Text.Body)
		}
	})

	t.Run("Invalid signature blocks processing", func(t *testing.T) {
		err := svc.ProcessEvent(context.Background(), payload, "sha256=invalidhash")
		if err == nil {
			t.Fatalf("expected error due to invalid signature")
		}
	})
}
