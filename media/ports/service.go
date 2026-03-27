package ports

import (
	"context"
	"io"
)

// MediaService orchestrates uploads and secure binary downloads of attachments.
type MediaService interface {
	DownloadMedia(ctx context.Context, mediaID string) (io.ReadCloser, error)
	UploadMedia(ctx context.Context, fileReader io.Reader, mimeType string) (string, error)
}
