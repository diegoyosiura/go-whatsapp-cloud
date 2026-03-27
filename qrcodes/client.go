package qrcodes

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/services"
)

// Client is the public Facade for the QRCode Management API module.
type Client struct {
	service ports.QRCodeService
}

// NewClient initializes the QRCode subsystem with valid graph keys.
func NewClient(apiVersion, phoneNumberID, token string) *Client {
	httpAdapter := adapters.NewHTTPQRCodeAdapter(apiVersion, phoneNumberID, token)
	service := services.NewQRCodeService(httpAdapter)

	return &Client{
		service: service,
	}
}

// Get queries detailed properties from an existing QRCode ID string.
func (c *Client) Get(ctx context.Context, codeID string) (domain.QRCode, error) {
	return c.service.Get(ctx, codeID)
}

// List returns all active QRCodes deployed for this Phone Number.
func (c *Client) List(ctx context.Context) ([]domain.QRCode, error) {
	return c.service.List(ctx)
}

// Create generates a new Physical QR code mapping to a text template. Format is commonly 'SVG' or 'PNG'.
func (c *Client) Create(ctx context.Context, prefilledMessage, format string) (domain.QRCode, error) {
	return c.service.Create(ctx, prefilledMessage, format)
}

// Update edits the prefilled template of an otherwise already distributed QRCode ID.
func (c *Client) Update(ctx context.Context, codeID, prefilledMessage string) (domain.QRCode, error) {
	return c.service.Update(ctx, codeID, prefilledMessage)
}

// Delete gracefully burns an existing QRCode.
func (c *Client) Delete(ctx context.Context, codeID string) (bool, error) {
	return c.service.Delete(ctx, codeID)
}
