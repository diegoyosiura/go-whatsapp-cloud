package domain

import "encoding/json"

// MessageTemplate represents a message template registered in Meta Business Manager.
type MessageTemplate struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Language   string          `json:"language"`
	Status     string          `json:"status"`
	Category   string          `json:"category"`
	Components json.RawMessage `json:"components"`
}

// MessageTemplateList is the paginated response from Meta's message_templates endpoint.
type MessageTemplateList struct {
	Data   []MessageTemplate `json:"data"`
	Paging *Paging           `json:"paging"`
}

// Paging holds Meta's cursor-based pagination info.
type Paging struct {
	Next string `json:"next"`
}
