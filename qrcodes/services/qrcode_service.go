package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/ports"
)

type qrCodeService struct {
	client ports.QRCodeClient
}

// NewQRCodeService builds the business layer decoupled from actual HTTP requests.
func NewQRCodeService(client ports.QRCodeClient) ports.QRCodeService {
	return &qrCodeService{
		client: client,
	}
}

func (s *qrCodeService) Get(ctx context.Context, codeID string) (domain.QRCode, error) {
	return s.client.Get(ctx, codeID)
}

func (s *qrCodeService) List(ctx context.Context) ([]domain.QRCode, error) {
	return s.client.List(ctx)
}

func (s *qrCodeService) Create(ctx context.Context, prefilledMessage, format string) (domain.QRCode, error) {
	return s.client.Create(ctx, prefilledMessage, format)
}

func (s *qrCodeService) Update(ctx context.Context, codeID, prefilledMessage string) (domain.QRCode, error) {
	return s.client.Update(ctx, codeID, prefilledMessage)
}

func (s *qrCodeService) Delete(ctx context.Context, codeID string) (bool, error) {
	return s.client.Delete(ctx, codeID)
}
