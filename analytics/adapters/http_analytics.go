package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/analytics/ports"
)

type httpAnalyticsAdapter struct {
	apiVersion string
	token      string
	client     ports.HTTPDoer
}

// NewHTTPAnalyticsAdapter acts as the network caller for analytics querying.
func NewHTTPAnalyticsAdapter(apiVersion, token string) ports.AnalyticsClient {
	return &httpAnalyticsAdapter{
		apiVersion: apiVersion,
		token:      token,
		client:     &http.Client{},
	}
}

func formatSlice(s []string) string {
	if len(s) == 0 {
		return "[]"
	}
	quoted := make([]string, len(s))
	for i, v := range s {
		quoted[i] = fmt.Sprintf(`"%s"`, v)
	}
	return fmt.Sprintf("[%s]", strings.Join(quoted, ","))
}

func (a *httpAnalyticsAdapter) GetAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, phoneNumbers, countryCodes []string) (*domain.Analytics, error) {
	phones := formatSlice(phoneNumbers)
	countries := formatSlice(countryCodes)

	query := fmt.Sprintf("analytics.start(%d).end(%d).granularity(%s).phone_numbers(%s).country_codes(%s)", start, end, granularity, phones, countries)

	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s?fields=%s", a.apiVersion, wabaID, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var parsed domain.AnalyticsResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return parsed.Analytics, nil
}

func (a *httpAnalyticsAdapter) GetConversationAnalytics(ctx context.Context, wabaID string, start, end int, granularity string, dimensions, conversationDirections []string) (*domain.Analytics, error) {
	dims := formatSlice(dimensions)
	dirs := formatSlice(conversationDirections)

	query := fmt.Sprintf("conversation_analytics.start(%d).end(%d).granularity(%s).conversation_directions(%s).dimensions(%s)", start, end, granularity, dirs, dims)

	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s?fields=%s", a.apiVersion, wabaID, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var parsed domain.AnalyticsResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return parsed.ConversationAnalytics, nil
}
