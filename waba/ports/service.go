package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/waba/domain"
)

// WABAService exposes natively Meta's global WhatsApp Business settings retrieval.
type WABAService interface {
	GetAccountInfo(ctx context.Context) (domain.AccountInfo, error)
	ListMessageTemplates(ctx context.Context) ([]domain.MessageTemplate, error)
}
