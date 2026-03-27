package uploads

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/services"
)

// Client is the facade to manage Resumable Upload sessions in Meta Graph API.
type Client struct {
	service ports.UploadService
}

// NewClient returns the initialized Resumable Upload subsystem.
func NewClient(apiVersion, token string) *Client {
	adapter := adapters.NewHTTPUploadAdapter(apiVersion, token)
	service := services.NewUploadService(adapter)

	return &Client{
		service: service,
	}
}

// CreateSession initializes a new upload chunk process. Returns the Upload ID.
func (c *Client) CreateSession(ctx context.Context, fileLength int, fileType, fileName string) (*domain.CreateSessionResponse, error) {
	return c.service.CreateSession(ctx, fileLength, fileType, fileName)
}

// UploadData sends binary chunk data matching a specific offset for a given Upload ID.
func (c *Client) UploadData(ctx context.Context, uploadID string, fileOffset int, fileData []byte, mimeType string) (*domain.UploadDataResponse, error) {
	return c.service.UploadData(ctx, uploadID, fileOffset, fileData, mimeType)
}

// QueryStatus returns the current committed offset size of a pending Resumable Upload session.
func (c *Client) QueryStatus(ctx context.Context, uploadID string) (*domain.QueryStatusResponse, error) {
	return c.service.QueryStatus(ctx, uploadID)
}
