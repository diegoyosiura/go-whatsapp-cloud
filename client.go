package whatsapp

import (
	"errors"
	"sync"

	"github.com/diegoyosiura/go-whatsapp-cloud/internal/core/domain"
)

// Client is the main Facade for the WhatsApp SDK.
// It manages configurations for one or more Phone Numbers (Multi-Tenant architecture).
type Client struct {
	mu      sync.RWMutex
	configs map[string]domain.WhatsAppConfig
}

// NewClient constructs a new instance of the SDK Client with a primary configuration.
// It stores the config keyed by its PhoneNumberID.
func NewClient(config domain.WhatsAppConfig) *Client {
	c := &Client{
		configs: make(map[string]domain.WhatsAppConfig),
	}

	if config.PhoneNumberID != "" {
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
