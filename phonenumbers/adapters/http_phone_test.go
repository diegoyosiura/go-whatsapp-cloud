package adapters

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/domain"
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

func TestHTTPPhoneAdapter_RequestCode(t *testing.T) {
	adapter := &httpPhoneAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Dispatches Verification correctly", func(t *testing.T) {
		err := adapter.RequestCode(context.Background(), domain.RequestCodePayload{
			CodeMethod: "SMS",
			Language:   "pt_BR",
		})
		
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		if mockDoer.Req.Method != http.MethodPost {
			t.Errorf("Method expected POST")
		}

		if !strings.HasSuffix(mockDoer.Req.URL.Path, "request_code") {
			t.Errorf("url ending mismatch: %s", mockDoer.Req.URL.Path)
		}

		bodyBytes, _ := io.ReadAll(mockDoer.Req.Body)
		if !strings.Contains(string(bodyBytes), `"code_method":"SMS"`) || !strings.Contains(string(bodyBytes), `"language":"pt_BR"`) {
			t.Errorf("Payload construction mismatch: %s", string(bodyBytes))
		}
	})

	t.Run("Captures HTTP 400 errors", func(t *testing.T) {
		mockDoer.Response.StatusCode = 400
		mockDoer.Response.Body = io.NopCloser(bytes.NewBufferString(`{"error": {"message": "Invalid method"}}`))

		err := adapter.RequestCode(context.Background(), domain.RequestCodePayload{CodeMethod: "WRONG"})
		if err == nil {
			t.Errorf("Expected 400 rejection")
		} else if !strings.Contains(err.Error(), "Invalid method") {
			t.Errorf("Error string did not propagate JSON payload: %v", err)
		}
	})
}

func TestHTTPPhoneAdapter_VerifyCode(t *testing.T) {
	adapter := &httpPhoneAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Submits PIN cleanly", func(t *testing.T) {
		err := adapter.VerifyCode(context.Background(), domain.VerifyCodePayload{Code: "123456"})
		if err != nil {
			t.Fatalf("unexpected failure: %v", err)
		}

		bodyBytes, _ := io.ReadAll(mockDoer.Req.Body)
		if !strings.Contains(string(bodyBytes), `"code":"123456"`) {
			t.Errorf("Payload missing code: %s", string(bodyBytes))
		}
	})

	t.Run("Handles network errors on Verify", func(t *testing.T) {
		mockDoer.Err = errors.New("timeout")
		err := adapter.VerifyCode(context.Background(), domain.VerifyCodePayload{Code: "111"})
		if err == nil {
			t.Errorf("expected timeout")
		}
	})
}

func TestHTTPPhoneAdapter_SetTwoStepVerification(t *testing.T) {
	adapter := &httpPhoneAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Submits PIN cleanly", func(t *testing.T) {
		err := adapter.SetTwoStepVerification(context.Background(), domain.SetTwoStepVerificationPayload{Pin: "123456"})
		if err != nil {
			t.Fatalf("unexpected failure: %v", err)
		}

		bodyBytes, _ := io.ReadAll(mockDoer.Req.Body)
		if !strings.Contains(string(bodyBytes), `"pin":"123456"`) {
			t.Errorf("Payload missing pin: %s", string(bodyBytes))
		}
	})
}

func TestHTTPPhoneAdapter_BlockUsers(t *testing.T) {
	adapter := &httpPhoneAdapter{apiVersion: "v20.0", phoneNumberID: "PHONE_ID", token: "VALID_TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Blocks User", func(t *testing.T) {
		err := adapter.BlockUser(context.Background(), domain.BlockUserRequest{
			MessagingProduct: "whatsapp",
			BlockUsers: []domain.BlockUserDetail{{User: "123"}},
		})
		if err != nil {
			t.Fatalf("unexpected failure: %v", err)
		}
		if mockDoer.Req.Method != http.MethodPost {
			t.Errorf("Method expected POST")
		}
	})

	t.Run("Unblocks User", func(t *testing.T) {
		err := adapter.UnblockUser(context.Background(), domain.BlockUserRequest{
			MessagingProduct: "whatsapp",
			BlockUsers: []domain.BlockUserDetail{{User: "123"}},
		})
		if err != nil {
			t.Fatalf("unexpected failure: %v", err)
		}
		if mockDoer.Req.Method != http.MethodDelete {
			t.Errorf("Method expected DELETE")
		}
	})

	t.Run("Gets Blocked users", func(t *testing.T) {
		mockDoer.Response.Body = io.NopCloser(strings.NewReader(`{"data": [{"wa_id": "4444"}]}`))
		res, err := adapter.GetBlockedUsers(context.Background())
		if err != nil {
			t.Fatalf("unexpected failure: %v", err)
		}
		if mockDoer.Req.Method != http.MethodGet {
			t.Errorf("Method expected GET")
		}
		if len(res.Data) == 0 || res.Data[0].WAID != "4444" {
			t.Errorf("Mismatch")
		}
	})
}
