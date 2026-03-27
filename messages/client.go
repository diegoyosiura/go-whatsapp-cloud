package messages

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/adapters"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/ports"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/services"
)

// Client is the public Facade for the Messages Sender module.
type Client struct {
	service ports.MessageService
}

// NewClient initializes the Messages subsystem with necessary API keys.
// apiVersion: like "v20.0"
// phoneNumberID: the ID of the sending phone number from Meta
// token: the User Access Token for the Authorization Bearer
func NewClient(apiVersion, phoneNumberID, token string) *Client {
	httpSender := adapters.NewHTTPSender(apiVersion, phoneNumberID, token)
	msgService := services.NewMessageService(httpSender)

	return &Client{
		service: msgService,
	}
}

// SendText dispatches a simple text message.
func (c *Client) SendText(ctx context.Context, to string, text string) (domain.MessageResponse, error) {
	return c.service.SendText(ctx, to, text)
}

// SendTemplate dispatches an approved Meta Template Message.
func (c *Client) SendTemplate(ctx context.Context, to, templateName, languageCode string, components []domain.TemplateComponent) (domain.MessageResponse, error) {
	return c.service.SendTemplate(ctx, to, templateName, languageCode, components)
}

// SendImage dispatches an image message by providing a public URL or Meta Media ID.
func (c *Client) SendImage(ctx context.Context, to, imageStr string, isID bool) (domain.MessageResponse, error) {
	return c.service.SendImage(ctx, to, imageStr, isID)
}

// SendInteractiveButton dispatches a message containing quick reply buttons.
func (c *Client) SendInteractiveButton(ctx context.Context, to, body string, buttons []domain.Button) (domain.MessageResponse, error) {
	return c.service.SendInteractiveButton(ctx, to, body, buttons)
}
