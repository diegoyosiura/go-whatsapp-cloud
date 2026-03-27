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

// SetTwoStepVerification allows altering the security PIN tied to the Phone Number.
func (c *Client) SetTwoStepVerification(ctx context.Context, pin string) error {
	return c.service.SetTwoStepVerification(ctx, pin)
}

// BlockUser prevents a specific WhatsApp ID from routing messages back to the bot.
func (c *Client) BlockUser(ctx context.Context, waID string) error {
	return c.service.BlockUser(ctx, waID)
}

// UnblockUser reverses a previously established Block constraint.
func (c *Client) UnblockUser(ctx context.Context, waID string) error {
	return c.service.UnblockUser(ctx, waID)
}

// GetBlockedUsers retrieves a linear string array of blocked target IDs.
func (c *Client) GetBlockedUsers(ctx context.Context) ([]string, error) {
	return c.service.GetBlockedUsers(ctx)
}
