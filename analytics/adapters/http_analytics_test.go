package adapters

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPDoer struct {
	Response *http.Response
	Err      error
	Req      *http.Request
}

func (m *mockHTTPDoer) Do(req *http.Request) (*http.Response, error) {
	m.Req = req
	return m.Response, m.Err
}

func TestHTTPAnalyticsAdapter_GetAnalytics(t *testing.T) {
	adapter := &httpAnalyticsAdapter{apiVersion: "v20.0", token: "TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"analytics": {
					"granularity": "DAY",
					"data_points": [{"start": 1, "end": 2, "sent": 10}]
				},
				"id": "WABA123"
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Query mapping", func(t *testing.T) {
		res, err := adapter.GetAnalytics(context.Background(), "WABA123", 1, 2, "DAY", []string{"1650"}, []string{"BR", "US"})
		if err != nil {
			t.Fatalf("unexpected %v", err)
		}
		if res.Granularity != "DAY" {
			t.Errorf("wrong parsed data")
		}
		if len(res.DataPoints) != 1 {
			t.Errorf("wrong data points count")
		}
		
		query := mockDoer.Req.URL.Query().Get("fields")
		if !strings.Contains(query, "analytics.start(1).end(2).granularity(DAY).phone_numbers([%221650%22]).country_codes([%22BR%22,%22US%22])") && !strings.Contains(query, "analytics.start(1).end(2).granularity(DAY).phone_numbers([\"1650\"]).country_codes([\"BR\",\"US\"])") {
			t.Errorf("wrong fields query mapping: %s", query)
		}
	})
}

func TestHTTPAnalyticsAdapter_GetConversationAnalytics(t *testing.T) {
	adapter := &httpAnalyticsAdapter{apiVersion: "v20.0", token: "TOKEN"}
	mockDoer := &mockHTTPDoer{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"conversation_analytics": {
					"granularity": "MONTHLY"
				},
				"id": "WABA123"
			}`)),
		},
	}
	adapter.client = mockDoer

	t.Run("Conversation Query mapping", func(t *testing.T) {
		res, err := adapter.GetConversationAnalytics(context.Background(), "WABA123", 1, 2, "MONTHLY", []string{"type"}, []string{"outbound"})
		if err != nil {
			t.Fatalf("unexpected %v", err)
		}
		if res.Granularity != "MONTHLY" {
			t.Errorf("wrong parsed data")
		}
		
		query := mockDoer.Req.URL.Query().Get("fields")
		if !strings.Contains(query, "dimensions") || !strings.Contains(query, "conversation_directions") {
			t.Errorf("missing query components: %s", query)
		}
	})
}
