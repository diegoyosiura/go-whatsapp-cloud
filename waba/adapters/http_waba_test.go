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

func TestHTTPWabaAdapter_GetAccountInfo(t *testing.T) {
	adapter := &httpWABAAdapter{apiVersion: "v20.0", wabaID: "VALID_WABA", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"id": "waba_id_888",
				"name": "Global Corp",
				"currency": "BRL",
				"timezone_id": "America/Sao_Paulo",
				"message_template_namespace": "unique_namespace_guid"
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Fetches WABA Properties matching API Docs", func(t *testing.T) {
		info, err := adapter.GetAccountInfo(context.Background())
		if err != nil {
			t.Fatalf("unexpected mock error: %v", err)
		}

		if info.ID != "waba_id_888" || info.Name != "Global Corp" {
			t.Errorf("Decoded JSON is missing properties: %+v", info)
		}

		if mockDoer.Req.Method != http.MethodGet {
			t.Errorf("Expected GET request")
		}

		if !strings.Contains(mockDoer.Req.URL.Path, "VALID_WABA") {
			t.Errorf("URL Path is skipping the WABA ID constraint (%s)", mockDoer.Req.URL.Path)
		}

		if mockDoer.Req.Header.Get("Authorization") != "Bearer VALID_TOKEN" {
			t.Errorf("Authorization Bearer wasn't forwarded")
		}
	})

	t.Run("Handles HTTP Statuses Over 400", func(t *testing.T) {
		mockDoer.Response.StatusCode = 401
		mockDoer.Response.Body = io.NopCloser(bytes.NewBufferString(`{"error": {"message": "Invalid OAuth"}}`))

		_, err := adapter.GetAccountInfo(context.Background())
		if err == nil {
			t.Fatalf("Expected error to be parsed from Graph payload")
		}

		if !strings.Contains(err.Error(), "Invalid OAuth") {
			t.Errorf("Error doesn't contain message from meta API: %v", err)
		}
	})

	t.Run("Fails Hard on Network Disconnections", func(t *testing.T) {
		mockDoer.Err = errors.New("timeout dialing host")
		_, err := adapter.GetAccountInfo(context.Background())
		if err == nil {
			t.Errorf("Expected native timeout err")
		}
	})
}
