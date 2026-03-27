package services

import (
	"context"
	"errors"
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/domain"
)

type mockAnalyticsClient struct {
	Res *domain.Analytics
	Err error
}

func (m *mockAnalyticsClient) GetAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, phoneNumbers, countryCodes []string) (*domain.Analytics, error) {
	return m.Res, m.Err
}

func (m *mockAnalyticsClient) GetConversationAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, dimensions, conversationDirections []string) (*domain.Analytics, error) {
	return m.Res, m.Err
}

func TestAnalyticsService(t *testing.T) {
	client := &mockAnalyticsClient{}
	s := NewAnalyticsService(client)

	t.Run("GetAnalytics success", func(t *testing.T) {
		client.Res = &domain.Analytics{Granularity: "DAY"}
		client.Err = nil
		res, err := s.GetAnalytics(context.Background(), "waba", 1, 2, "DAY", nil, nil)
		if err != nil {
			t.Fatalf("unexpected %v", err)
		}
		if res.Granularity != "DAY" {
			t.Errorf("mismatch granularity")
		}
	})

	t.Run("Throws error", func(t *testing.T) {
		client.Err = errors.New("fail")
		_, err := s.GetConversationAnalytics(context.Background(), "waba", 1, 2, "DAY", nil, nil)
		if err == nil {
			t.Errorf("expected error")
		}
	})
}
