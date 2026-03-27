package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles/ports"
)

type httpProfileAdapter struct {
	apiVersion    string
	phoneNumberID string
	token         string
	client        ports.HTTPDoer
}

func NewHTTPProfileAdapter(apiVersion, phoneNumberID, token string) ports.ProfileClient {
	return &httpProfileAdapter{
		apiVersion:    apiVersion,
		phoneNumberID: phoneNumberID,
		token:         token,
		client:        &http.Client{},
	}
}

// GetProfile fetche the currently active WhatsApp Business Profile settings directly from Meta API.
func (a *httpProfileAdapter) GetProfile(ctx context.Context) (domain.ProfileResponse, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/whatsapp_business_profile?fields=about,address,description,email,profile_picture_url,websites,vertical", a.apiVersion, a.phoneNumberID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.ProfileResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return domain.ProfileResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return domain.ProfileResponse{}, fmt.Errorf("api returned status: %d", resp.StatusCode)
	}

	var parsed domain.ProfileResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return domain.ProfileResponse{}, err
	}

	return parsed, nil
}

// UpdateProfile mutations the Business Profile settings using the fields established by the Meta Graph.
func (a *httpProfileAdapter) UpdateProfile(ctx context.Context, profile domain.ProfileUpdate) error {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/whatsapp_business_profile", a.apiVersion, a.phoneNumberID)

	// Since we use Initialisms rule for domain, JSON structure implicitly mirrors Graph API keys via standard marshal.
	bodyBytes, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("failed packing payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		// Reads snippet of the raw error for traceability.
		var errData struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if defErr := json.NewDecoder(resp.Body).Decode(&errData); defErr != nil {
			return fmt.Errorf("api declined status %d and bad body: %v", resp.StatusCode, defErr)
		}
		return fmt.Errorf("api declined status %d: %s", resp.StatusCode, errData.Error.Message)
	}

	return nil
}
