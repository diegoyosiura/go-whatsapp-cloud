package whatsapp

import (
	"errors"
	"sync"

	"github.com/diegoyosiura/go-whatsapp-cloud/analytics"
	"github.com/diegoyosiura/go-whatsapp-cloud/businessprofiles"
	"github.com/diegoyosiura/go-whatsapp-cloud/internal/core/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/media"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages"
	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers"
	"github.com/diegoyosiura/go-whatsapp-cloud/qrcodes"
	"github.com/diegoyosiura/go-whatsapp-cloud/uploads"
	"github.com/diegoyosiura/go-whatsapp-cloud/waba"
	"github.com/diegoyosiura/go-whatsapp-cloud/webhook"
)

// Client is the main Facade for the WhatsApp SDK.
// It manages configurations for one or more Phone Numbers (Multi-Tenant architecture).
type Client struct {
	mu      sync.RWMutex
	primary string
	configs map[string]domain.WhatsAppConfig
}

// NewClient constructs a new instance of the SDK Client with a primary configuration.
// It stores the config keyed by its PhoneNumberID.
func NewClient(config domain.WhatsAppConfig) *Client {
	c := &Client{
		configs: make(map[string]domain.WhatsAppConfig),
	}

	if config.PhoneNumberID != "" {
		c.primary = config.PhoneNumberID
		c.configs[config.PhoneNumberID] = config
	}

	return c
}

// AddPhoneNumberConfig registers an additional WhatsApp configuration dynamically.
// This is safe for concurrent access.
func (c *Client) AddPhoneNumberConfig(id string, config domain.WhatsAppConfig) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.configs[id] = config
}

// GetPhoneNumberConfig retrieves the configuration bound to a specific phone number ID.
// Returns an error if the configuration is not registered.
func (c *Client) GetPhoneNumberConfig(id string) (domain.WhatsAppConfig, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	config, exists := c.configs[id]
	if !exists {
		return domain.WhatsAppConfig{}, errors.New("configuration not found for the given phone number id")
	}

	return config, nil
}

// PhoneContext isolates the API scope to a specific Phone Number inside a Multi-Tenant setup.
type PhoneContext struct {
	config domain.WhatsAppConfig
}

// ForPhone scopes the SDK operations to a specific registered phone number ID.
func (c *Client) ForPhone(phoneNumberID string) (*PhoneContext, error) {
	cfg, err := c.GetPhoneNumberConfig(phoneNumberID)
	if err != nil {
		return nil, err
	}
	return &PhoneContext{config: cfg}, nil
}

// Primary routes automatically to the first config injected in NewClient.
func (c *Client) Primary() *PhoneContext {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return &PhoneContext{config: c.configs[c.primary]}
}

// Webhook initializes the Webhook Listener and Validator module.
// It is globally scoped to the App and does not require a specific Phone Number.
func (c *Client) Webhook(appSecret, verifyToken string) *webhook.Client {
	return webhook.NewClient(appSecret, verifyToken)
}

// Messages returns the Facade to send text, media, and template messages.
func (p *PhoneContext) Messages() *messages.Client {
	return messages.NewClient(p.config.Version, p.config.PhoneNumberID, p.config.UserAccessToken)
}

// Media returns the Facade to manage image, video, and document blobs.
func (p *PhoneContext) Media() *media.Client {
	return media.NewClient(p.config.Version, p.config.PhoneNumberID, p.config.UserAccessToken)
}

// BusinessProfiles connects to the endpoints to read and update the Account Profile.
func (p *PhoneContext) BusinessProfiles() *businessprofiles.Client {
	return businessprofiles.NewClient(p.config.Version, p.config.PhoneNumberID, p.config.UserAccessToken)
}

// PhoneNumbers configures advanced features like Spam Blocking and 2FA.
func (p *PhoneContext) PhoneNumbers() *phonenumbers.Client {
	return phonenumbers.NewClient(p.config.Version, p.config.PhoneNumberID, p.config.UserAccessToken)
}

// QRCodes constructs the module to manage WhatsApp QR Code Generation.
func (p *PhoneContext) QRCodes() *qrcodes.Client {
	return qrcodes.NewClient(p.config.Version, p.config.PhoneNumberID, p.config.UserAccessToken)
}

// WABA provides Account-level configurations (Currency, Node Namespaces). Relies on WabaID.
func (p *PhoneContext) WABA() *waba.Client {
	return waba.NewClient(p.config.Version, p.config.WABAID, p.config.UserAccessToken)
}

// Analytics retrieves conversation and message billing throughput metrics. Relies on WabaID.
func (p *PhoneContext) Analytics() *analytics.Client {
	return analytics.NewClient(p.config.Version, p.config.UserAccessToken)
}

// Uploads exposes Resumable Upload sessions useful for +100MB files.
func (p *PhoneContext) Uploads() *uploads.Client {
	return uploads.NewClient(p.config.Version, p.config.UserAccessToken)
}
