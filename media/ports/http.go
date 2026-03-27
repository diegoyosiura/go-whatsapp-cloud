package ports

import (
	"context"
	"io"
	"net/http"
	
	"github.com/diegoyosiura/go-whatsapp-cloud/media/domain"
)

// HTTPDoer abstracts the standard net/http Client for TDD mockability.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// MediaClient abstracts the Meta APIs specifically for Media retrieval and submission.
type MediaClient interface {
	GetMediaInfo(ctx context.Context, mediaID string) (domain.MediaInfo, error)
	DownloadBinary(ctx context.Context, mediaURL string) (io.ReadCloser, error)
	UploadBinary(ctx context.Context, fileReader io.Reader, mimeType string) (string, error)
}
