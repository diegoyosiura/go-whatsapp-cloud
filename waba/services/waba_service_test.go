package services

import (
	"context"
	"errors"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/waba/domain"
)

type mockWabaClient struct {
	Response domain.AccountInfo
	Err      error
}

func (m *mockWabaClient) GetAccountInfo(ctx context.Context) (domain.AccountInfo, error) {
	return m.Response, m.Err
}

func TestWabaService_GetAccountInfo(t *testing.T) {
	mockClient := &mockWabaClient{}
	service := NewWABAService(mockClient)

	t.Run("Fetches Account Info Smoothly", func(t *testing.T) {
		mockClient.Err = nil
		mockClient.Response = domain.AccountInfo{
			ID:       "waba123",
			Name:     "Test Inc.",
			Currency: "USD",
		}

		info, err := service.GetAccountInfo(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if info.ID != "waba123" || info.Currency != "USD" {
			t.Errorf("Mismatch in payload: %+v", info)
		}
	})

	t.Run("Forwards Errors Intact", func(t *testing.T) {
		mockClient.Err = errors.New("bad request")
		_, err := service.GetAccountInfo(context.Background())
		if err == nil {
			t.Errorf("Expected bad request err")
		}
	})
}
