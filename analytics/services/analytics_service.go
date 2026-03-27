package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/ports"
)

type analyticsService struct {
	client ports.AnalyticsClient
}

// NewAnalyticsService creates the analytics usecase orchestrator.
func NewAnalyticsService(client ports.AnalyticsClient) ports.AnalyticsService {
	return &analyticsService{client: client}
}

func (s *analyticsService) GetAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, phoneNumbers, countryCodes []string) (*domain.Analytics, error) {
	return s.client.GetAnalytics(ctx, wabaID, start, end, granularity, phoneNumbers, countryCodes)
}

func (s *analyticsService) GetConversationAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, dimensions, conversationDirections []string) (*domain.Analytics, error) {
	return s.client.GetConversationAnalytics(ctx, wabaID, start, end, granularity, dimensions, conversationDirections)
}
