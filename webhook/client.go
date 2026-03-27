package webhook

import (
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/webhook/services"
)

// Client is the public Facade for the Webhook module.
// It hides the internal complexity of Orchestrators, Validators and Parsers.
type Client struct {
	processor ports.WebhookProcessor
	handler   http.Handler
}

// NewClient initializes the Webhook subsystem with the required secrets/tokens.
// secret: The App Secret used to validate X-Hub-Signature-256 for POST events.
// verifyToken: The custom token used by Meta to verify the GET webhook endpoint.
func NewClient(secret, verifyToken string) *Client {
	processor := services.NewWebhookService(secret)
	httpHandler := adapters.NewWebhookHTTPHandler(processor, verifyToken)

	return &Client{
		processor: processor,
		handler:   httpHandler,
	}
}

// RegisterEventHandler attaches an implementation of ports.EventHandler 
// to listen to incoming metadata events like Messages or Statuses.
func (c *Client) RegisterEventHandler(handler ports.EventHandler) {
	c.processor.RegisterHandler(handler)
}

// HTTPHandler exposes the standard net/http handler adapter 
// to be easily plugged into any Go standard HTTP server or framework.
func (c *Client) HTTPHandler() http.Handler {
	return c.handler
}
