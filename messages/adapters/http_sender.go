package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/ports"
)

// httpSender handles real HTTP dispatching to WhatsApp/Meta Graph API.
type httpSender struct {
	apiVersion    string
	phoneNumberID string
	token         string
	doer          ports.HTTPDoer
}

// NewHTTPSender initializes the adapter with Meta's endpoint info.
func NewHTTPSender(apiVersion, phoneNumberID, token string) ports.MessageSender {
	return &httpSender{
		apiVersion:    apiVersion,
		phoneNumberID: phoneNumberID,
		token:         token,
		doer:          &http.Client{},
	}
}

// SendMessage dispatches the payload to the Cloud API.
func (s *httpSender) SendMessage(ctx context.Context, payload domain.SendMessagePayload) (domain.MessageResponse, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", s.apiVersion, s.phoneNumberID)

	body, err := json.Marshal(payload)
	if err != nil {
		return domain.MessageResponse{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return domain.MessageResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.doer.Do(req)
	if err != nil {
		return domain.MessageResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		// Log detailed error from Meta in real conditions, for now simply returning status
		return domain.MessageResponse{}, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var messageResponse domain.MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
		return domain.MessageResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return messageResponse, nil
}
