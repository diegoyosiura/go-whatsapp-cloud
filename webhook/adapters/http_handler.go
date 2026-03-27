package adapters

import (
	"io"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/ports"
)

// webhookHTTPHandler is the internal adapter for go's net/http.
type webhookHTTPHandler struct {
	processor   ports.WebhookProcessor
	verifyToken string
}

// NewWebhookHTTPHandler creates a new net/http compatible handler interface
func NewWebhookHTTPHandler(processor ports.WebhookProcessor, verifyToken string) http.Handler {
	return &webhookHTTPHandler{
		processor:   processor,
		verifyToken: verifyToken,
	}
}

// ServeHTTP implements the http.Handler interface natively
func (h *webhookHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Meta Webhook Verification
		mode := r.URL.Query().Get("hub.mode")
		token := r.URL.Query().Get("hub.verify_token")
		challenge := r.URL.Query().Get("hub.challenge")

		if mode == "subscribe" && token == h.verifyToken {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(challenge))
			return
		}

		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method == http.MethodPost {
		// Event Processing
		signature := r.Header.Get("X-Hub-Signature-256")
		
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// The orchestrator handles signature validation and deep parsing.
		// WhatsApp expects a 200 OK fast response before deep processing, 
		// but for simplicity and correct TDD we process it synchronously here. 
		// Real-world scenarios might decouple this into a goroutine.
		err = h.processor.ProcessEvent(r.Context(), body, signature)
		if err != nil {
			// Even on processor failure (bad signature, invalid payload), 
			// it's a good practice to return 200 to Meta unless we specifically want them to retry.
			// But since we want to prevent bad actors, we could return 400.
			// In our test, we assert 200 OK regardless to acknowledge receipt.
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
