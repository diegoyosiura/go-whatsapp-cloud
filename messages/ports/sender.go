package ports

import (
	"context"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
)

// MessageSender defines the contract required to dispatch any kind of message to Meta.
type MessageSender interface {
	SendMessage(ctx context.Context, payload domain.SendMessagePayload) (domain.MessageResponse, error)
}

// HTTPDoer is an interface defining the minimal contract of net/http.Client
// It helps in mocking standard library behavior during tests.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
