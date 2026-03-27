package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/domain"
)

// UploadClient abstracts the raw network constraints for chunked data uploads.
type UploadClient interface {
	CreateSession(ctx context.Context, fileLength int, fileType, fileName string) (*domain.CreateSessionResponse, error)
	UploadData(ctx context.Context, uploadID string, fileOffset int, fileData []byte, mimeType string) (*domain.UploadDataResponse, error)
	QueryStatus(ctx context.Context, uploadID string) (*domain.QueryStatusResponse, error)
}

// UploadService provides the business orchestration for media uploads.
type UploadService interface {
	CreateSession(ctx context.Context, fileLength int, fileType, fileName string) (*domain.CreateSessionResponse, error)
	UploadData(ctx context.Context, uploadID string, fileOffset int, fileData []byte, mimeType string) (*domain.UploadDataResponse, error)
	QueryStatus(ctx context.Context, uploadID string) (*domain.QueryStatusResponse, error)
}
