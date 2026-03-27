package media

import (
	"context"
	"io"

	"github.com/diegoyosiura/go-whatsapp-cloud/media/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/media/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/media/services"
)

// Client is the public Facade for the Media Management module.
type Client struct {
	service ports.MediaService
}

// NewClient initializes the Media subsystem.
func NewClient(apiVersion, phoneNumberID, token string) *Client {
	httpAdapter := adapters.NewHTTPMediaAdapter(apiVersion, phoneNumberID, token)
	mediaService := services.NewMediaService(httpAdapter)

	return &Client{
		service: mediaService,
	}
}

// DownloadMedia fetches the encrypted Meta URL for the given mediaID and downloads the binary payload.
// The caller is responsible for eventually closing the returned io.ReadCloser to avoid memory leaks.
func (c *Client) DownloadMedia(ctx context.Context, mediaID string) (io.ReadCloser, error) {
	return c.service.DownloadMedia(ctx, mediaID)
}

// UploadMedia streams a binary payload from an io.Reader to WhatsApp Cloud API via multipart form-data.
// Returns the newly generated Media ID (wamid) on success.
func (c *Client) UploadMedia(ctx context.Context, fileReader io.Reader, mimeType string) (string, error) {
	return c.service.UploadMedia(ctx, fileReader, mimeType)
}
