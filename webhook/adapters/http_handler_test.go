package adapters

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/ports"
)

type mockProcessor struct {
	payload   []byte
	signature string
}

func (m *mockProcessor) ProcessEvent(ctx context.Context, payload []byte, signature string) error {
	m.payload = payload
	m.signature = signature
	return nil
}

func (m *mockProcessor) RegisterHandler(handler ports.EventHandler) {}

func TestWebhookHTTPHandler(t *testing.T) {
	processor := &mockProcessor{}
	verifyToken := "my_verify_token"

	handler := NewWebhookHTTPHandler(processor, verifyToken)

	t.Run("GET - Successful Verification", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/webhook?hub.mode=subscribe&hub.challenge=112233&hub.verify_token="+verifyToken, nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected 200 OK, got %d", rec.Code)
		}

		if rec.Body.String() != "112233" {
			t.Errorf("expected challenge 112233, got %s", rec.Body.String())
		}
	})

	t.Run("GET - Failed Verification", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/webhook?hub.mode=subscribe&hub.challenge=112233&hub.verify_token=wrong_token", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusForbidden {
			t.Errorf("expected 403 Forbidden, got %d", rec.Code)
		}
	})

	t.Run("POST - Event Processing", func(t *testing.T) {
		body := `{"test":"payload"}`
		req := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(body))
		req.Header.Set("X-Hub-Signature-256", "sha256=1234")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected 200 OK, got %d", rec.Code)
		}

		if string(processor.payload) != body {
			t.Errorf("expected payload %s, got %s", body, string(processor.payload))
		}

		if processor.signature != "sha256=1234" {
			t.Errorf("expected signature 'sha256=1234', got '%s'", processor.signature)
		}
	})
}
