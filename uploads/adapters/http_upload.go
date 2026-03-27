package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/uploads/ports"
)

type httpUploadAdapter struct {
	apiVersion string
	token      string
	client     ports.HTTPDoer
}

// NewHTTPUploadAdapter allocates an adapter specifically for Meta's Resumable Upload specification.
func NewHTTPUploadAdapter(apiVersion, token string) ports.UploadClient {
	return &httpUploadAdapter{
		apiVersion: apiVersion,
		token:      token,
		client:     &http.Client{},
	}
}

func (a *httpUploadAdapter) CreateSession(ctx context.Context, fileLength int, fileType, fileName string) (*domain.CreateSessionResponse, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/app/uploads", a.apiVersion)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("file_length", strconv.Itoa(fileLength))
	q.Add("file_type", fileType)
	if fileName != "" {
		q.Add("file_name", fileName)
	}
	req.URL.RawQuery = q.Encode()

	// Special Note: CreateSession explicitly requires OAuth prefix
	req.Header.Set("Authorization", "OAuth "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var parsed domain.CreateSessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}

func (a *httpUploadAdapter) UploadData(ctx context.Context, uploadID string, fileOffset int, fileData []byte, mimeType string) (*domain.UploadDataResponse, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s", a.apiVersion, uploadID)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(fileData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", mimeType)
	req.Header.Set("file_offset", strconv.Itoa(fileOffset))
	req.Header.Set("Authorization", "OAuth "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var parsed domain.UploadDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}

func (a *httpUploadAdapter) QueryStatus(ctx context.Context, uploadID string) (*domain.QueryStatusResponse, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s", a.apiVersion, uploadID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "OAuth "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var parsed domain.QueryStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}
