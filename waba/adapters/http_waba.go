package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/waba/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/waba/ports"
)

// httpWABAAdapter translates pure struct interactions into standard Meta URL requests.
type httpWABAAdapter struct {
	apiVersion string
	wabaID     string
	token      string
	client     ports.HTTPDoer
}

// NewHTTPWABAAdapter instantiates the network bridge passing the 'WABA ID' as parameter instead of 'Phone ID'.
func NewHTTPWABAAdapter(apiVersion, wabaID, token string) ports.WABAClient {
	return &httpWABAAdapter{
		apiVersion: apiVersion,
		wabaID:     wabaID,
		token:      token,
		client:     &http.Client{},
	}
}

// GetAccountInfo calls the global settings edge for the Waba Root node.
func (a *httpWABAAdapter) GetAccountInfo(ctx context.Context) (domain.AccountInfo, error) {
	// The Graph API for WABA specifies fields querystring to retrieve global metrics.
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s?fields=id,name,currency,timezone_id,message_template_namespace", a.apiVersion, a.wabaID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.AccountInfo{}, err
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

	resp, err := a.client.Do(req)
	if err != nil {
		return domain.AccountInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errData struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if defErr := json.NewDecoder(resp.Body).Decode(&errData); defErr != nil {
			return domain.AccountInfo{}, fmt.Errorf("api declined status %d and bad body: %v", resp.StatusCode, defErr)
		}
		return domain.AccountInfo{}, fmt.Errorf("api declined status %d: %s", resp.StatusCode, errData.Error.Message)
	}

	var parsed domain.AccountInfo
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return domain.AccountInfo{}, err
	}

	return parsed, nil
}
