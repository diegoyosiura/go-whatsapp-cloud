package ports

import (
	"context"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
)

// HTTPDoer acts as an abstraction over the native net/http Client roundtripper.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// ProfileClient establishes the pure hexagonal contract.
type ProfileClient interface {
	UpdateProfile(ctx context.Context, profile domain.ProfileUpdate) error
	GetProfile(ctx context.Context) (domain.ProfileResponse, error)
}
