package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
)

// ProfileService exposes cleanly the internal manipulations of business profiles.
type ProfileService interface {
	UpdateProfile(ctx context.Context, about, address, email string, websites []string) error
	GetProfile(ctx context.Context) (domain.ProfileResponse, error)
}
