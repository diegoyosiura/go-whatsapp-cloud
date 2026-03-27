package ports

import (
	"context"
	"net/http"

	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/domain"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type PhoneClient interface {
	RequestCode(ctx context.Context, payload domain.RequestCodePayload) error
	VerifyCode(ctx context.Context, payload domain.VerifyCodePayload) error
}
