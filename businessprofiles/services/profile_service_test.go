package services

import (
	"context"
	"errors"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
)

type mockProfileClient struct {
	UpdatedProfile domain.ProfileUpdate
	UpdateErr      error
	GetResponse    domain.ProfileResponse
	GetErr         error
}

func (m *mockProfileClient) UpdateProfile(ctx context.Context, profile domain.ProfileUpdate) error {
	m.UpdatedProfile = profile
	return m.UpdateErr
}

func (m *mockProfileClient) GetProfile(ctx context.Context) (domain.ProfileResponse, error) {
	return m.GetResponse, m.GetErr
}

func TestProfileService_Update(t *testing.T) {
	mockClient := &mockProfileClient{}
	service := NewProfileService(mockClient)

	t.Run("Valid Update Forwards to Client", func(t *testing.T) {
		mockClient.UpdateErr = nil

		// Testing builder via service wrapper
		err := service.UpdateProfile(context.Background(), "New About", "123 Street", "company@mail.com", []string{"https://x.com"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if mockClient.UpdatedProfile.About != "New About" ||
			mockClient.UpdatedProfile.Address != "123 Street" ||
			mockClient.UpdatedProfile.Email != "company@mail.com" ||
			len(mockClient.UpdatedProfile.Websites) != 1 {
			t.Errorf("profile parameters did not propagate correctly: %+v", mockClient.UpdatedProfile)
		}
	})

	t.Run("Fails gracefully when API rejects", func(t *testing.T) {
		mockClient.UpdateErr = errors.New("invalid website url")
		err := service.UpdateProfile(context.Background(), "", "", "", []string{"htp/invalid"})
		
		if err == nil || err.Error() != "invalid website url" {
			t.Fatalf("expected specific error, got %v", err)
		}
	})
}
