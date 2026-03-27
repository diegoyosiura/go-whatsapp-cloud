package adapters

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
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

func TestHTTPProfileAdapter_GetProfile(t *testing.T) {
	adapter := &httpProfileAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": [
					{
						"about": "Welcome to my shop",
						"address": "NY",
						"email": "a@b.com",
						"websites": ["https://foo.bar"]
					}
				]
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Get Profile successfully decodes JSON", func(t *testing.T) {
		resp, err := adapter.GetProfile(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(resp.Data) == 0 {
			t.Fatalf("expected data array")
		}

		if resp.Data[0].About != "Welcome to my shop" {
			t.Errorf("expected welcome message")
		}

		if mockDoer.Req.Header.Get("Authorization") != "Bearer VALID_TOKEN" {
			t.Errorf("missing bearer token")
		}
	})

	t.Run("Fails gracefully on Network error", func(t *testing.T) {
		mockDoer.Err = errors.New("timeout")
		_, err := adapter.GetProfile(context.Background())
		if err == nil {
			t.Errorf("expected error on timeout")
		}
	})
}

func TestHTTPProfileAdapter_UpdateProfile(t *testing.T) {
	adapter := &httpProfileAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Sends POST request correctly with JSON", func(t *testing.T) {
		err := adapter.UpdateProfile(context.Background(), domain.ProfileUpdate{About: "New Info"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if mockDoer.Req.Method != http.MethodPost {
			t.Errorf("expected POST method")
		}

		bodyBytes, _ := io.ReadAll(mockDoer.Req.Body)
		if !strings.Contains(string(bodyBytes), `"about":"New Info"`) {
			t.Errorf("JSON did not securely serialize the struct, got: %s", string(bodyBytes))
		}
		
		if mockDoer.Req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("mismatched content type")
		}
	})
}
