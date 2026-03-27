package businessprofiles

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/services"
)

// Client is the entry point for manipulating the WhatsApp Business Account Profile. 
// It adheres closely to Hexagonal Architecture, instantiating services and HTTP adapters seamlessly.
type Client struct {
	service ports.ProfileService
}

// NewClient returns an instance bound to the caller's Graph API credentials.
func NewClient(apiVersion, phoneNumberID, token string) *Client {
	adapter := adapters.NewHTTPProfileAdapter(apiVersion, phoneNumberID, token)
	service := services.NewProfileService(adapter)

	return &Client{
		service: service,
	}
}

// UpdateProfile commits immediate changes to the business profile info of the target phone number.
func (c *Client) UpdateProfile(ctx context.Context, about, address, email string, websites []string) error {
	return c.service.UpdateProfile(ctx, about, address, email, websites)
}

// GetProfile fetches the live Business Profile properties registered under this Phone Number ID.
func (c *Client) GetProfile(ctx context.Context) (domain.ProfileResponse, error) {
	return c.service.GetProfile(ctx)
}
