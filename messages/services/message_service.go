package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/usecases"
)

type messageService struct {
	sender ports.MessageSender
}

// NewMessageService instantiates the high level service to send diverse messages
func NewMessageService(sender ports.MessageSender) ports.MessageService {
	return &messageService{
		sender: sender,
	}
}

// SendText simplifies sending a text message using the underlying builder and adapter.
func (s *messageService) SendText(ctx context.Context, to string, body string) (domain.MessageResponse, error) {
	payload := usecases.BuildTextPayload(to, body)
	return s.sender.SendMessage(ctx, payload)
}

// SendTemplate simplifies sending a WhatsApp approved template.
func (s *messageService) SendTemplate(ctx context.Context, to, templateName, languageCode string, components []domain.TemplateComponent) (domain.MessageResponse, error) {
	payload := usecases.BuildTemplatePayload(to, templateName, languageCode, components)
	return s.sender.SendMessage(ctx, payload)
}

// SendImage simplifies sending an image via URL or Media ID.
func (s *messageService) SendImage(ctx context.Context, to, imageStr string, isID bool) (domain.MessageResponse, error) {
	payload := usecases.BuildImagePayload(to, imageStr, isID)
	return s.sender.SendMessage(ctx, payload)
}

// SendInteractiveButton simplifies sending interactive buttons.
func (s *messageService) SendInteractiveButton(ctx context.Context, to, body string, buttons []domain.Button) (domain.MessageResponse, error) {
	payload := usecases.BuildInteractiveButtonPayload(to, body, buttons)
	return s.sender.SendMessage(ctx, payload)
}
