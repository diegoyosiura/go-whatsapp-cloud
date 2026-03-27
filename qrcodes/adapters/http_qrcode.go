package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes/ports"
)

type httpQRCodeAdapter struct {
	apiVersion    string
	phoneNumberID string
	token         string
	client        ports.HTTPDoer
}

// NewHTTPQRCodeAdapter parses Graph configuration and exposes pure struct operations.
func NewHTTPQRCodeAdapter(apiVersion, phoneNumberID, token string) ports.QRCodeClient {
	return &httpQRCodeAdapter{
		apiVersion:    apiVersion,
		phoneNumberID: phoneNumberID,
		token:         token,
		client:        &http.Client{},
	}
}

func (a *httpQRCodeAdapter) Get(ctx context.Context, codeID string) (domain.QRCode, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/message_qrdls/%s", a.apiVersion, a.phoneNumberID, codeID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.QRCode{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return domain.QRCode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return domain.QRCode{}, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var parsed domain.QRCodeListResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return domain.QRCode{}, err
	}
	if len(parsed.Data) == 0 {
		return domain.QRCode{}, fmt.Errorf("no code found")
	}

	return parsed.Data[0], nil
}

func (a *httpQRCodeAdapter) List(ctx context.Context) ([]domain.QRCode, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/message_qrdls", a.apiVersion, a.phoneNumberID)

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

	var parsed domain.QRCodeListResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return parsed.Data, nil
}

func (a *httpQRCodeAdapter) Create(ctx context.Context, prefilledMessage, format string) (domain.QRCode, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/message_qrdls", a.apiVersion, a.phoneNumberID)

	payload := domain.CreateQRCodeRequest{
		PrefilledMessage: prefilledMessage,
		GenerateQRImage:  format,
	}
	bodyBytes, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return domain.QRCode{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return domain.QRCode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return domain.QRCode{}, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var code domain.QRCode
	if err := json.NewDecoder(resp.Body).Decode(&code); err != nil {
		return domain.QRCode{}, err
	}

	return code, nil
}

func (a *httpQRCodeAdapter) Update(ctx context.Context, codeID, prefilledMessage string) (domain.QRCode, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/message_qrdls", a.apiVersion, a.phoneNumberID)

	payload := domain.UpdateQRCodeRequest{
		Code:             codeID,
		PrefilledMessage: prefilledMessage,
	}
	bodyBytes, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return domain.QRCode{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return domain.QRCode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return domain.QRCode{}, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var code domain.QRCode
	if err := json.NewDecoder(resp.Body).Decode(&code); err != nil {
		return domain.QRCode{}, err
	}

	return code, nil
}

func (a *httpQRCodeAdapter) Delete(ctx context.Context, codeID string) (bool, error) {
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/message_qrdls/%s", a.apiVersion, a.phoneNumberID, codeID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return false, fmt.Errorf("api declined status %d", resp.StatusCode)
	}

	var successRes domain.SuccessResponse
	if err := json.NewDecoder(resp.Body).Decode(&successRes); err != nil {
		return false, err
	}

	return successRes.Success, nil
}
