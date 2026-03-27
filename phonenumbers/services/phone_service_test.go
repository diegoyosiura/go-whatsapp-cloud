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
	SetTwoStepPayload domain.SetTwoStepVerificationPayload
	BlockPayload      domain.BlockUserRequest
	UnblockPayload    domain.BlockUserRequest
	BlockedUsersRes   *domain.GetBlockedUsersResponse
	GenericErr        error
}

func (m *mockPhoneClient) RequestCode(ctx context.Context, payload domain.RequestCodePayload) error {
	m.ReqPayload = payload
	return m.ReqErr
}

func (m *mockPhoneClient) VerifyCode(ctx context.Context, payload domain.VerifyCodePayload) error {
	m.VerifyPayload = payload
	return m.VerifyErr
}

func (m *mockPhoneClient) SetTwoStepVerification(ctx context.Context, payload domain.SetTwoStepVerificationPayload) error {
	m.SetTwoStepPayload = payload
	return m.GenericErr
}

func (m *mockPhoneClient) BlockUser(ctx context.Context, payload domain.BlockUserRequest) error {
	m.BlockPayload = payload
	return m.GenericErr
}

func (m *mockPhoneClient) UnblockUser(ctx context.Context, payload domain.BlockUserRequest) error {
	m.UnblockPayload = payload
	return m.GenericErr
}

func (m *mockPhoneClient) GetBlockedUsers(ctx context.Context) (*domain.GetBlockedUsersResponse, error) {
	return m.BlockedUsersRes, m.GenericErr
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

func TestPhoneService_SetTwoStepVerification(t *testing.T) {
	mockClient := &mockPhoneClient{}
	service := NewPhoneService(mockClient)

	t.Run("Sets PIN", func(t *testing.T) {
		mockClient.GenericErr = nil
		err := service.SetTwoStepVerification(context.Background(), "123456")
		if err != nil {
			t.Fatalf("unexpected nil error")
		}
		if mockClient.SetTwoStepPayload.Pin != "123456" {
			t.Errorf("PIN mismatch")
		}
	})
}

func TestPhoneService_BlockManagement(t *testing.T) {
	mockClient := &mockPhoneClient{}
	service := NewPhoneService(mockClient)

	t.Run("Blocks user", func(t *testing.T) {
		mockClient.GenericErr = nil
		err := service.BlockUser(context.Background(), "4444")
		if err != nil {
			t.Fatalf("unexpected err")
		}
		if len(mockClient.BlockPayload.BlockUsers) == 0 || mockClient.BlockPayload.BlockUsers[0].User != "4444" {
			t.Errorf("user not mapped")
		}
	})

	t.Run("Unblocks user", func(t *testing.T) {
		mockClient.GenericErr = nil
		err := service.UnblockUser(context.Background(), "5555")
		if err != nil {
			t.Fatalf("unexpected err")
		}
		if len(mockClient.UnblockPayload.BlockUsers) == 0 || mockClient.UnblockPayload.BlockUsers[0].User != "5555" {
			t.Errorf("user not mapped")
		}
	})

	t.Run("Gets blocked users", func(t *testing.T) {
		mockClient.GenericErr = nil
		mockClient.BlockedUsersRes = &domain.GetBlockedUsersResponse{
			Data: []domain.GetBlockedUsersData{
				{WAID: "6666"},
			},
		}
		users, err := service.GetBlockedUsers(context.Background())
		if err != nil {
			t.Fatalf("unexpected err")
		}
		if len(users) == 0 || users[0] != "6666" {
			t.Errorf("failed extraction")
		}
	})
}
