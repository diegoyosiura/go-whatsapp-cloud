package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/domain"
)

// QRCodeClient defines the Network HTTP capabilities to Graph.
type QRCodeClient interface {
	Get(ctx context.Context, codeID string) (domain.QRCode, error)
	List(ctx context.Context) ([]domain.QRCode, error)
	Create(ctx context.Context, prefilledMessage, format string) (domain.QRCode, error)
	Update(ctx context.Context, codeID, prefilledMessage string) (domain.QRCode, error)
	Delete(ctx context.Context, codeID string) (bool, error)
}

// QRCodeService acts as the orchestrator hiding REST HTTP context.
type QRCodeService interface {
	Get(ctx context.Context, codeID string) (domain.QRCode, error)
	List(ctx context.Context) ([]domain.QRCode, error)
	Create(ctx context.Context, prefilledMessage, format string) (domain.QRCode, error)
	Update(ctx context.Context, codeID, prefilledMessage string) (domain.QRCode, error)
	Delete(ctx context.Context, codeID string) (bool, error)
}
