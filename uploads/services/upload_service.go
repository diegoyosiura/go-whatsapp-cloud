package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/ports"
)

type uploadService struct {
	client ports.UploadClient
}

// NewUploadService creates the isolation logic wrapping the raw HTTP calls.
func NewUploadService(client ports.UploadClient) ports.UploadService {
	return &uploadService{client: client}
}

func (s *uploadService) CreateSession(ctx context.Context, fileLength int, fileType, fileName string) (*domain.CreateSessionResponse, error) {
	return s.client.CreateSession(ctx, fileLength, fileType, fileName)
}

func (s *uploadService) UploadData(ctx context.Context, uploadID string, fileOffset int, fileData []byte, mimeType string) (*domain.UploadDataResponse, error) {
	return s.client.UploadData(ctx, uploadID, fileOffset, fileData, mimeType)
}

func (s *uploadService) QueryStatus(ctx context.Context, uploadID string) (*domain.QueryStatusResponse, error) {
	return s.client.QueryStatus(ctx, uploadID)
}
