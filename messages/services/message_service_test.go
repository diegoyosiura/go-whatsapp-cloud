package services

import (
	"context"
	"errors"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
)

type mockMessageSender struct {
	Response domain.MessageResponse
	Err      error
	Payload  domain.SendMessagePayload
}

func (m *mockMessageSender) SendMessage(ctx context.Context, payload domain.SendMessagePayload) (domain.MessageResponse, error) {
	m.Payload = payload
	return m.Response, m.Err
}

func TestMessageService_SendText(t *testing.T) {
	mockSender := &mockMessageSender{
		Response: domain.MessageResponse{
			Messages: []domain.ResponseMessage{{ID: "wamid.123"}},
		},
	}

	service := NewMessageService(mockSender)

	t.Run("Successfully send text", func(t *testing.T) {
		resp, err := service.SendText(context.Background(), "551199999", "Hello TDD")

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.Messages[0].ID != "wamid.123" {
			t.Errorf("expected wamid.123, got %s", resp.Messages[0].ID)
		}

		if mockSender.Payload.Text.Body != "Hello TDD" {
			t.Errorf("expected payload to have text 'Hello TDD', got '%s'", mockSender.Payload.Text.Body)
		}
	})

	t.Run("Sender Error Propagation", func(t *testing.T) {
		mockSender.Err = errors.New("network timeout")
		
		_, err := service.SendText(context.Background(), "551199999", "Should Fail")

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
