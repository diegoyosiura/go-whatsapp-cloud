package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/usecases"
)

type webhookService struct {
	secret   string
	handlers []ports.EventHandler
}

// NewWebhookService returns a fresh WebhookProcessor implementation.
func NewWebhookService(secret string) ports.WebhookProcessor {
	return &webhookService{
		secret:   secret,
		handlers: make([]ports.EventHandler, 0),
	}
}

func (s *webhookService) RegisterHandler(handler ports.EventHandler) {
	s.handlers = append(s.handlers, handler)
}

func (s *webhookService) ProcessEvent(ctx context.Context, payload []byte, signature string) error {
	// 1. Validate signature
	err := usecases.ValidateSignature(payload, signature, s.secret)
	if err != nil {
		return err
	}

	// 2. Parse payload
	parsedPayload, err := usecases.ParsePayload(payload)
	if err != nil {
		return err
	}

	// 3. Route to handlers
	for _, entry := range parsedPayload.Entry {
		for _, change := range entry.Changes {
			if change.Field == "messages" {
				// Process messages
				for _, msg := range change.Value.Messages {
					for _, handler := range s.handlers {
						_ = handler.OnMessage(ctx, msg) // ignoring errors for now (could log them)
					}
				}
				// Process statuses
				for _, status := range change.Value.Statuses {
					for _, handler := range s.handlers {
						_ = handler.OnStatus(ctx, status)
					}
				}
			}
		}
	}

	return nil
}
