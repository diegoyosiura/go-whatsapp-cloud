package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/media/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/media/ports"
)

type httpMediaAdapter struct {
	apiVersion    string
	phoneNumberID string
	token         string
	client        ports.HTTPDoer
}

func NewHTTPMediaAdapter(apiVersion, phoneNumberID, token string) ports.MediaClient {
	return &httpMediaAdapter{
		apiVersion:    apiVersion,
		phoneNumberID: phoneNumberID,
		token:         token,
		client:        &http.Client{},
	}
}

// GetMediaInfo requests the absolute tracking URL from a Media ID.
func (a *httpMediaAdapter) GetMediaInfo(ctx context.Context, mediaID string) (domain.MediaInfo, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s", a.apiVersion, mediaID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.MediaInfo{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return domain.MediaInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return domain.MediaInfo{}, fmt.Errorf("api returned status: %d", resp.StatusCode)
	}

	var info domain.MediaInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return domain.MediaInfo{}, err
	}
	return info, nil
}

// DownloadBinary issues a GET request to the encrypted URL containing the Meta stream.
func (a *httpMediaAdapter) DownloadBinary(ctx context.Context, mediaURL string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mediaURL, nil)
	if err != nil {
		return nil, err
	}
	// Meta requires the bearer token to download the encrypted blob.
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("api returned status: %d when downloading binary", resp.StatusCode)
	}

	// Deliberately NOT closing resp.Body because we yield it to the caller (io.ReadCloser).
	return resp.Body, nil
}

// UploadBinary builds a multipart/form-data request encoding the binary and streams to Meta.
func (a *httpMediaAdapter) UploadBinary(ctx context.Context, fileReader io.Reader, mimeType string) (string, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/media", a.apiVersion, a.phoneNumberID)

	// Since Go's explicit http multipart requires a complete buffer in memory using bytes.Buffer,
	// this approach is okay for minor docs/images up to typical limits.
	var bodyBuf bytes.Buffer
	writer := multipart.NewWriter(&bodyBuf)

	if err := writer.WriteField("messaging_product", "whatsapp"); err != nil {
		return "", err
	}
	
	// 'file' field accepts the payload
	// Setting type directly is slightly tricky natively without setting headers per sub-part:
	partHeaders := make(map[string][]string)
	partHeaders["Content-Disposition"] = []string{`form-data; name="file"; filename="upload"`}
	partHeaders["Content-Type"] = []string{mimeType}
	
	part, err := writer.CreatePart(partHeaders)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(part, fileReader); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &bodyBuf)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("api returned status %d on upload", resp.StatusCode)
	}

	var upResp domain.UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&upResp); err != nil {
		return "", err
	}

	return upResp.ID, nil
}
