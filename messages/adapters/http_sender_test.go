package adapters

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
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

func TestHTTPSender_SendMessage(t *testing.T) {
	apiVersion := "v20.0"
	phoneNumberID := "PHONE123"
	token := "valid_token"
	sender := &httpSender{
		apiVersion:    apiVersion,
		phoneNumberID: phoneNumberID,
		token:         token,
	}

	payload := domain.SendMessagePayload{
		MessagingProduct: "whatsapp",
		To:               "123",
		Type:             "text",
		Text: &domain.TextObject{
			Body: "Hello",
		},
	}

	t.Run("Successful send", func(t *testing.T) {
		mockDoer := &mockHTTPDoer{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"messaging_product": "whatsapp",
					"messages": [{"id": "wamid.123"}]
				}`)),
			},
		}

		// Injecting the mock client
		sender.doer = mockDoer

		resp, err := sender.SendMessage(context.Background(), payload)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(resp.Messages) == 0 || resp.Messages[0].ID != "wamid.123" {
			t.Errorf("expected wamid.123, got %v", resp)
		}

		// Check injected headers and URL
		if mockDoer.Req.URL.String() != "https://graph.facebook.com/v20.0/PHONE123/messages" {
			t.Errorf("unexpected URL: %s", mockDoer.Req.URL.String())
		}
		if mockDoer.Req.Header.Get("Authorization") != "Bearer valid_token" {
			t.Errorf("missing or invalid authorization header")
		}
		if mockDoer.Req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("missing content-type header")
		}
	})

	t.Run("API Error Handling", func(t *testing.T) {
		mockDoer := &mockHTTPDoer{
			Response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"error": {"message": "Invalid parameter"}
				}`)),
			},
		}
		sender.doer = mockDoer

		_, err := sender.SendMessage(context.Background(), payload)

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		expectedError := "API returned status: 400"
		if err.Error() != expectedError {
			t.Errorf("expected '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("Network Error", func(t *testing.T) {
		mockDoer := &mockHTTPDoer{
			Err: errors.New("timeout"),
		}
		sender.doer = mockDoer

		_, err := sender.SendMessage(context.Background(), payload)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
