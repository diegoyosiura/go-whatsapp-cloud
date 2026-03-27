package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/ports"
)

// httpPhoneAdapter connects the abstract service to real Meta API Web Endpoints.
type httpPhoneAdapter struct {
	apiVersion    string
	phoneNumberID string
	token         string
	client        ports.HTTPDoer
}

// NewHTTPPhoneAdapter initializes the networking constraints required by Meta.
func NewHTTPPhoneAdapter(apiVersion, phoneNumberID, token string) ports.PhoneClient {
	return &httpPhoneAdapter{
		apiVersion:    apiVersion,
		phoneNumberID: phoneNumberID,
		token:         token,
		client:        &http.Client{},
	}
}

func (a *httpPhoneAdapter) RequestCode(ctx context.Context, payload domain.RequestCodePayload) error {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/request_code", a.apiVersion, a.phoneNumberID)
	return a.doPost(ctx, url, payload)
}

func (a *httpPhoneAdapter) VerifyCode(ctx context.Context, payload domain.VerifyCodePayload) error {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/verify_code", a.apiVersion, a.phoneNumberID)
	return a.doPost(ctx, url, payload)
}

func (a *httpPhoneAdapter) SetTwoStepVerification(ctx context.Context, payload domain.SetTwoStepVerificationPayload) error {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s", a.apiVersion, a.phoneNumberID)
	return a.doPost(ctx, url, payload)
}

func (a *httpPhoneAdapter) BlockUser(ctx context.Context, payload domain.BlockUserRequest) error {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/block_users", a.apiVersion, a.phoneNumberID)
	return a.doPost(ctx, url, payload)
}

func (a *httpPhoneAdapter) UnblockUser(ctx context.Context, payload domain.BlockUserRequest) error {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/block_users", a.apiVersion, a.phoneNumberID)
	bodyBytes, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewBuffer(bodyBytes))
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
		return fmt.Errorf("api declined status %d", resp.StatusCode)
	}
	return nil
}

func (a *httpPhoneAdapter) GetBlockedUsers(ctx context.Context) (*domain.GetBlockedUsersResponse, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/block_users", a.apiVersion, a.phoneNumberID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
	var res domain.GetBlockedUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// doPost is a private helper that DRYs the networking logic.
func (a *httpPhoneAdapter) doPost(ctx context.Context, url string, payload interface{}) error {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed unpacking json payload: %w", err)
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
		var errData struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		json.NewDecoder(resp.Body).Decode(&errData)
		return fmt.Errorf("api declined status %d: %s", resp.StatusCode, errData.Error.Message)
	}

	return nil
}
