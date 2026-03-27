package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/ports"
)

type profileService struct {
	client ports.ProfileClient
}

// NewProfileService initializes the service responsible for business profile interactions.
func NewProfileService(client ports.ProfileClient) ports.ProfileService {
	return &profileService{
		client: client,
	}
}

// UpdateProfile is a helper that constructs the payload struct and delegates to the HTTP client port.
func (s *profileService) UpdateProfile(ctx context.Context, about, address, email string, websites []string) error {
	payload := domain.ProfileUpdate{
		About:    about,
		Address:  address,
		Email:    email,
		Websites: websites,
	}

	return s.client.UpdateProfile(ctx, payload)
}

// GetProfile delegates the fetch call logic cleanly.
func (s *profileService) GetProfile(ctx context.Context) (domain.ProfileResponse, error) {
	return s.client.GetProfile(ctx)
}
