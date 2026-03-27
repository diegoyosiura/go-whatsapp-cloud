package services

import (
	"context"
	"errors"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/domain"
)

type mockPhoneClient struct {
	ReqPayload    domain.RequestCodePayload
	ReqErr        error
	VerifyPayload domain.VerifyCodePayload
	VerifyErr     error
}

func (m *mockPhoneClient) RequestCode(ctx context.Context, payload domain.RequestCodePayload) error {
	m.ReqPayload = payload
	return m.ReqErr
}

func (m *mockPhoneClient) VerifyCode(ctx context.Context, payload domain.VerifyCodePayload) error {
	m.VerifyPayload = payload
	return m.VerifyErr
}

func TestPhoneService_RequestCode(t *testing.T) {
	mockClient := &mockPhoneClient{}
	service := NewPhoneService(mockClient)

	t.Run("Successfully builds the SMS request snippet", func(t *testing.T) {
		mockClient.ReqErr = nil
		err := service.RequestCode(context.Background(), "SMS", "pt_BR")
		
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if mockClient.ReqPayload.CodeMethod != "SMS" {
			t.Errorf("Code Method mismatch")
		}

		if mockClient.ReqPayload.Language != "pt_BR" {
			t.Errorf("Language mismatch")
		}
	})

	t.Run("Captures errors generated globally", func(t *testing.T) {
		mockClient.ReqErr = errors.New("rate limit")
		err := service.RequestCode(context.Background(), "VOICE", "en")
		if err == nil {
			t.Errorf("expected global rate limit error")
		}
	})
}

func TestPhoneService_VerifyCode(t *testing.T) {
	mockClient := &mockPhoneClient{}
	service := NewPhoneService(mockClient)

	t.Run("Sends verification cleanly", func(t *testing.T) {
		mockClient.VerifyErr = nil
		err := service.VerifyCode(context.Background(), "234567")

		if err != nil {
			t.Fatalf("expected nil error")
		}

		if mockClient.VerifyPayload.Code != "234567" {
			t.Errorf("Code mismatch")
		}
	})
}
