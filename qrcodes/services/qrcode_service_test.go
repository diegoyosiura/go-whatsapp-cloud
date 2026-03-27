package services

import (
	"context"
	"errors"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/domain"
)

type mockQRCodeClient struct {
	Response domain.QRCode
	ListRes  []domain.QRCode
	BoolRes  bool
	Err      error
}

func (m *mockQRCodeClient) Get(ctx context.Context, codeID string) (domain.QRCode, error) {
	return m.Response, m.Err
}
func (m *mockQRCodeClient) List(ctx context.Context) ([]domain.QRCode, error) {
	return m.ListRes, m.Err
}
func (m *mockQRCodeClient) Create(ctx context.Context, prefilledMessage, format string) (domain.QRCode, error) {
	return m.Response, m.Err
}
func (m *mockQRCodeClient) Update(ctx context.Context, codeID, prefilledMessage string) (domain.QRCode, error) {
	return m.Response, m.Err
}
func (m *mockQRCodeClient) Delete(ctx context.Context, codeID string) (bool, error) {
	return m.BoolRes, m.Err
}

func TestQRCodeService_Get(t *testing.T) {
	mockClient := &mockQRCodeClient{}
	service := NewQRCodeService(mockClient)

	t.Run("Successfully forwards GET", func(t *testing.T) {
		mockClient.Err = nil
		mockClient.Response = domain.QRCode{Code: "QR_123", PrefilledMessage: "Scan me"}

		res, err := service.Get(context.Background(), "QR_123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res.Code != "QR_123" {
			t.Errorf("mismatch payload")
		}
	})

	t.Run("Passes through errors", func(t *testing.T) {
		mockClient.Err = errors.New("timeout")
		_, err := service.Get(context.Background(), "QR")
		if err == nil {
			t.Errorf("expected error")
		}
	})
}
