package waba

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/waba/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/waba/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/waba/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/waba/services"
)

// Client handles all global account interactions directly targeting a WABA ID.
// Unlike `messages` or `phonenumbers`, this relies on WabaID parameter.
type Client struct {
	service ports.WABAService
}

// NewClient bootstraps the Administrative endpoints context.
func NewClient(apiVersion, wabaID, token string) *Client {
	adapter := adapters.NewHTTPWABAAdapter(apiVersion, wabaID, token)
	service := services.NewWABAService(adapter)

	return &Client{
		service: service,
	}
}

// GetAccountInfo extracts high-level Business information (Currency, Name, Node Namespace) from Graph API.
func (c *Client) GetAccountInfo(ctx context.Context) (domain.AccountInfo, error) {
	return c.service.GetAccountInfo(ctx)
}

// ListMessageTemplates fetches all message templates registered for this WABA.
func (c *Client) ListMessageTemplates(ctx context.Context) ([]domain.MessageTemplate, error) {
	return c.service.ListMessageTemplates(ctx)
}
