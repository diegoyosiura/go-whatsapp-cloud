package ports

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/domain"
)

// AnalyticsClient defines the HTTP contract to interact with Meta Graph.
type AnalyticsClient interface {
	GetAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, phoneNumbers, countryCodes []string) (*domain.Analytics, error)
	GetConversationAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, dimensions, conversationDirections []string) (*domain.Analytics, error)
}

// AnalyticsService defines the business orchestrator to fetch WABA analytics.
type AnalyticsService interface {
	GetAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, phoneNumbers, countryCodes []string) (*domain.Analytics, error)
	GetConversationAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, dimensions, conversationDirections []string) (*domain.Analytics, error)
}
