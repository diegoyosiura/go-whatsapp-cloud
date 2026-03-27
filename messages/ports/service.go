package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
)

type MessageService interface {
	SendText(ctx context.Context, to, text string) (domain.MessageResponse, error)
	SendTemplate(ctx context.Context, to, templateName, languageCode string, components []domain.TemplateComponent) (domain.MessageResponse, error)
	SendImage(ctx context.Context, to, imageStr string, isID bool) (domain.MessageResponse, error)
	SendInteractiveButton(ctx context.Context, to, body string, buttons []domain.Button) (domain.MessageResponse, error)
}
