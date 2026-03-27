package adapters

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPDoer struct {
	Response *http.Response
	Err      error
	Req      *http.Request
}

func (m *mockHTTPDoer) Do(req *http.Request) (*http.Response, error) {
	m.Req = req
	return m.Response, m.Err
}

func TestHTTPMediaAdapter_GetMediaInfo(t *testing.T) {
	adapter := &httpMediaAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"url": "https://secure.meta/image123",
				"mime_type": "image/jpeg",
				"sha256": "123hash",
				"file_size": 2048,
				"id": "media123"
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Successfully get info", func(t *testing.T) {
		info, err := adapter.GetMediaInfo(context.Background(), "media123")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if info.URL != "https://secure.meta/image123" {
			t.Errorf("url mismatch: %s", info.URL)
		}
		if mockDoer.Req.Header.Get("Authorization") != "Bearer VALID_TOKEN" {
			t.Errorf("missing auth header")
		}
	})

	t.Run("API Error Handling", func(t *testing.T) {
		mockDoer.Response.StatusCode = 404
		_, err := adapter.GetMediaInfo(context.Background(), "invalid")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestHTTPMediaAdapter_DownloadBinary(t *testing.T) {
	adapter := &httpMediaAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("binary content")),
		},
	}
	adapter.client = mockDoer

	t.Run("Download Successfully", func(t *testing.T) {
		stream, err := adapter.DownloadBinary(context.Background(), "https://secure.meta/image")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		data, _ := io.ReadAll(stream)
		if string(data) != "binary content" {
			t.Errorf("content mismatch")
		}
		// ensure auth token is sent on binary requests
		if mockDoer.Req.Header.Get("Authorization") != "Bearer VALID_TOKEN" {
			t.Errorf("missing auth token on binary request")
		}
	})
}

func TestHTTPMediaAdapter_UploadBinary(t *testing.T) {
	adapter := &httpMediaAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"h": "wamid.456"}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Upload Form Data", func(t *testing.T) {
		id, err := adapter.UploadBinary(context.Background(), strings.NewReader("fake image"), "image/jpeg")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if id != "wamid.456" {
			t.Errorf("expected wamid.456, got %s", id)
		}

		contentType := mockDoer.Req.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "multipart/form-data") {
			t.Errorf("invalid content-type for upload: %s", contentType)
		}

		// Ensure body was encoded as multipart
		reqBody, _ := io.ReadAll(mockDoer.Req.Body)
		if !strings.Contains(string(reqBody), `name="file"`) {
			t.Errorf("multipart missing 'file' field")
		}
		if !strings.Contains(string(reqBody), `name="messaging_product"`) {
			t.Errorf("multipart missing 'messaging_product' field")
		}
	})

	t.Run("Network Failure", func(t *testing.T) {
		mockDoer.Err = errors.New("timeout")
		_, err := adapter.UploadBinary(context.Background(), strings.NewReader("fake"), "image/jpeg")
		if err == nil {
			t.Errorf("expected network error")
		}
	})
}
