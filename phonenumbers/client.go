package phonenumbers

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/services"
)

// Client encapsulates the logic required to request and verify OTP tokens with the Cloud API
// configuring the WABA phone number.
type Client struct {
	service ports.PhoneService
}

// NewClient returns the initialized gateway for administrative phone configurations.
func NewClient(apiVersion, phoneNumberID, token string) *Client {
	adapter := adapters.NewHTTPPhoneAdapter(apiVersion, phoneNumberID, token)
	service := services.NewPhoneService(adapter)

	return &Client{
		service: service,
	}
}

// RequestVerificationCode triggers Meta to send a 6-digit PIN via SMS or VOICE.
func (c *Client) RequestVerificationCode(ctx context.Context, codeMethod, language string) error {
	return c.service.RequestCode(ctx, codeMethod, language)
}

// VerifyCode submits the 6-digit PIN back to Meta to finish the onboarding setup.
func (c *Client) VerifyCode(ctx context.Context, code string) error {
	return c.service.VerifyCode(ctx, code)
}
