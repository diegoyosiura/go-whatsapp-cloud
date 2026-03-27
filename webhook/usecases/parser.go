package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/domain"
)

// ParsePayload takes the raw JSON bytes from the Meta Webhook and serializes
// it into the struct domain.WebhookPayload. It returns an error if the JSON is malformed.
func ParsePayload(data []byte) (*domain.WebhookPayload, error) {
	var payload domain.WebhookPayload

	err := json.Unmarshal(data, &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse webhook payload: %w", err)
	}

	return &payload, nil
}
