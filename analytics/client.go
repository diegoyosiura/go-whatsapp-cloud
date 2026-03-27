package analytics

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/services"
)

// Client is the public Facade for the Analytics Extractor.
type Client struct {
	service ports.AnalyticsService
}

// NewClient initializes the Analytics subsystem pointing exclusively to the WABA Graph.
func NewClient(apiVersion, token string) *Client {
	httpAdapter := adapters.NewHTTPAnalyticsAdapter(apiVersion, token)
	service := services.NewAnalyticsService(httpAdapter)

	return &Client{
		service: service,
	}
}

// GetAnalytics fetches generic WABA messaging metrics.
func (c *Client) GetAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, phoneNumbers, countryCodes []string) (*domain.Analytics, error) {
	return c.service.GetAnalytics(ctx, wabaID, start, end, granularity, phoneNumbers, countryCodes)
}

// GetConversationAnalytics retrieves in-depth metrics grouping by conversation attributes.
func (c *Client) GetConversationAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, dimensions, conversationDirections []string) (*domain.Analytics, error) {
	return c.service.GetConversationAnalytics(ctx, wabaID, start, end, granularity, dimensions, conversationDirections)
}
