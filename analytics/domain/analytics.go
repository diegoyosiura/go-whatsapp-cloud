package domain

// DataPoint represents a single chunk of analytics data scoped to a time window.
type DataPoint struct {
	Start                int     `json:"start"`
	End                  int     `json:"end"`
	Sent                 int     `json:"sent,omitempty"`
	Delivered            int     `json:"delivered,omitempty"`
	Conversation         int     `json:"conversation,omitempty"`
	Cost                 float64 `json:"cost,omitempty"`
	ConversationType     string  `json:"conversation_type,omitempty"`
	ConversationDirection string `json:"conversation_direction,omitempty"`
}

// Analytics represents the common metric object for both messaging and conversations.
type Analytics struct {
	PhoneNumbers           []string    `json:"phone_numbers,omitempty"`
	CountryCodes           []string    `json:"country_codes,omitempty"`
	ConversationDirections []string    `json:"conversation_directions,omitempty"`
	Dimensions             []string    `json:"dimensions,omitempty"`
	Granularity            string      `json:"granularity,omitempty"`
	DataPoints             []DataPoint `json:"data_points,omitempty"`
}

// AnalyticsResponse encapsulates the Graph API WABA level analytics query result.
type AnalyticsResponse struct {
	Analytics             *Analytics `json:"analytics,omitempty"`
	ConversationAnalytics *Analytics `json:"conversation_analytics,omitempty"`
	ID                    string     `json:"id"`
}
