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

func TestHTTPQRCodeAdapter_Get(t *testing.T) {
	adapter := &httpQRCodeAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE", token: "TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": [
					{
						"code": "QR123",
						"prefilled_message": "Hello!",
						"deep_link_url": "wa.me"
					}
				]
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Parses specific code properly", func(t *testing.T) {
		qr, err := adapter.Get(context.Background(), "QR123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if qr.Code != "QR123" {
			t.Errorf("parsed wrongly: %s", qr.Code)
		}

		if mockDoer.Req.Method != http.MethodGet {
			t.Errorf("wrong method")
		}
		if !strings.Contains(mockDoer.Req.URL.Path, "PHONE/message_qrdls/QR123") {
			t.Errorf("path missing id %s", mockDoer.Req.URL.Path)
		}
	})
}

func TestHTTPQRCodeAdapter_Create(t *testing.T) {
	adapter := &httpQRCodeAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE", token: "TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"code": "QR123",
				"prefilled_message": "New"
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("POST request formatting ok", func(t *testing.T) {
		_, err := adapter.Create(context.Background(), "New", "SVG")
		if err != nil {
			t.Fatalf("unexpected %v", err)
		}

		if mockDoer.Req.Method != http.MethodPost {
			t.Errorf("wrong method")
		}
        
        reqBody, _ := io.ReadAll(mockDoer.Req.Body)
        if !strings.Contains(string(reqBody), "generate_qr_image") {
            t.Errorf("missing generate field")
        }
	})
}

func TestHTTPQRCodeAdapter_Update(t *testing.T) {
	adapter := &httpQRCodeAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE", token: "TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"code": "QR123",
				"prefilled_message": "Updated"
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("POST request passes the code", func(t *testing.T) {
		_, err := adapter.Update(context.Background(), "QR123", "Updated")
		if err != nil {
			t.Fatalf("unexpected %v", err)
		}
        
        reqBody, _ := io.ReadAll(mockDoer.Req.Body)
        if !strings.Contains(string(reqBody), `"code":"QR123"`) {
            t.Errorf("missing code definition")
        }
	})
}

func TestHTTPQRCodeAdapter_Delete(t *testing.T) {
	adapter := &httpQRCodeAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE", token: "TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"success": true
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("DELETE successfully processed", func(t *testing.T) {
		suc, err := adapter.Delete(context.Background(), "QR123")
		if err != nil {
			t.Fatalf("unexpected %v", err)
		}
        if !suc {
            t.Errorf("expected success true")
        }
		if mockDoer.Req.Method != http.MethodDelete {
			t.Errorf("wrong method")
		}
	})
}
