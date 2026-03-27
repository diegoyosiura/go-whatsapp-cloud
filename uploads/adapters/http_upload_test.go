package adapters

import (
	"bytes"
	"context"
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

func TestHTTPUploadAdapter_CreateSession(t *testing.T) {
	adapter := &httpUploadAdapter{apiVersion: "v20.0", token: "valid_token"}
	mockServer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"id": "session_id_123"}`)),
		},
	}
	adapter.client = mockServer

	res, err := adapter.CreateSession(context.Background(), 1024, "image/jpeg", "avatar.jpg")
	if err != nil {
		t.Fatalf("unexpected fail: %v", err)
	}
	if res.ID != "session_id_123" {
		t.Errorf("invalid mapping")
	}

	if mockServer.Req.Method != http.MethodPost {
		t.Errorf("expected POST method")
	}

	query := mockServer.Req.URL.Query()
	if query.Get("file_length") != "1024" || query.Get("file_type") != "image/jpeg" || query.Get("file_name") != "avatar.jpg" {
		t.Errorf("params invalid: %v", query)
	}
}

func TestHTTPUploadAdapter_UploadData(t *testing.T) {
	adapter := &httpUploadAdapter{apiVersion: "v20.0", token: "valid_token"}
	mockServer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"h": "handle_x1"}`)),
		},
	}
	adapter.client = mockServer

	res, err := adapter.UploadData(context.Background(), "session_id", 0, []byte("raw_binary"), "image/jpeg")
	if err != nil {
		t.Fatalf("unexpected fail: %v", err)
	}
	if res.Handle != "handle_x1" {
		t.Errorf("invalid mapping")
	}

	if mockServer.Req.Header.Get("file_offset") != "0" {
		t.Errorf("missing header mapping")
	}
	if !strings.HasPrefix(mockServer.Req.Header.Get("Authorization"), "OAuth") {
		t.Errorf("Resumable requires OAuth Prefix! %v", mockServer.Req.Header.Get("Authorization"))
	}
	body, _ := io.ReadAll(mockServer.Req.Body)
	if string(body) != "raw_binary" {
		t.Errorf("Binary payload mismatch")
	}
}

func TestHTTPUploadAdapter_QueryStatus(t *testing.T) {
	adapter := &httpUploadAdapter{apiVersion: "v20.0", token: "valid_token"}
	mockServer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"id": "session", "file_offset": 512}`)),
		},
	}
	adapter.client = mockServer

	res, err := adapter.QueryStatus(context.Background(), "session")
	if err != nil {
		t.Fatalf("unexpected err")
	}

	if res.FileOffset != 512 {
		t.Errorf("JSON mapping offset fail")
	}
}
