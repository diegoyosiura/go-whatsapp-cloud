package services

import (
	"context"
	"fmt"
	"io"

	"github.com/diegoyosiura/go-whatsapp-cloud/media/ports"
)

type mediaService struct {
	client ports.MediaClient
}

// NewMediaService initializes the media orchestrator.
func NewMediaService(client ports.MediaClient) ports.MediaService {
	return &mediaService{
		client: client,
	}
}

// DownloadMedia handles the 2-step process to get a binary from Meta:
// First it retrieves the encrypted URL endpoint, then it downloads the blob.
func (s *mediaService) DownloadMedia(ctx context.Context, mediaID string) (io.ReadCloser, error) {
	info, err := s.client.GetMediaInfo(ctx, mediaID)
	if err != nil {
		return nil, fmt.Errorf("failed fetching media info: %w", err)
	}

	binaryStream, err := s.client.DownloadBinary(ctx, info.URL)
	if err != nil {
		return nil, fmt.Errorf("failed fetching binary stream: %w", err)
	}

	return binaryStream, nil
}

// UploadMedia relays the stream context to the Media Client, bypassing any intermediate caching.
func (s *mediaService) UploadMedia(ctx context.Context, fileReader io.Reader, mimeType string) (string, error) {
	id, err := s.client.UploadBinary(ctx, fileReader, mimeType)
	if err != nil {
		return "", fmt.Errorf("failed to upload media: %w", err)
	}
	return id, nil
}
