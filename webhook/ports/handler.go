package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/domain"
)

// EventHandler defines the contract for consuming incoming WhatsApp events.
// Developers using the SDK can implement this interface to process incoming messages, statuses, etc.
type EventHandler interface {
	OnMessage(ctx context.Context, message domain.Message) error
	OnStatus(ctx context.Context, status domain.Status) error
}

// WebhookProcessor represents the internal orchestrator that processes the raw HTTP request,
// validates its signature, parses the payload and dispatches events to registered EventHandlers.
type WebhookProcessor interface {
	ProcessEvent(ctx context.Context, payload []byte, signature string) error
	RegisterHandler(handler EventHandler)
}
