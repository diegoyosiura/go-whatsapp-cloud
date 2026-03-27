package services

import (
	"context"

	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/domain"
	"github.com/diegoyosiura/go-whatsapp-cloud/phonenumbers/ports"
)

// phoneService acts as an intermediator abstracting Meta constraints towards native calls.
type phoneService struct {
	client ports.PhoneClient
}

// NewPhoneService initializes the wrapper that binds business logic to the adapter.
func NewPhoneService(client ports.PhoneClient) ports.PhoneService {
	return &phoneService{
		client: client,
	}
}

// RequestCode commands Meta WhatsApp Cloud API to emit a verification PIN code (OTP) via SMS or Voice.
// Requires method ('SMS' | 'VOICE') and locale language, e.g., 'en', 'pt_BR'.
func (s *phoneService) RequestCode(ctx context.Context, codeMethod, language string) error {
	payload := domain.RequestCodePayload{
		CodeMethod: codeMethod,
		Language:   language,
	}

	return s.client.RequestCode(ctx, payload)
}

// VerifyCode confirms the 6-digits PIN sent by Meta's RequestCode call towards onboarding a Number.
func (s *phoneService) VerifyCode(ctx context.Context, code string) error {
	payload := domain.VerifyCodePayload{
		Code: code,
	}

	return s.client.VerifyCode(ctx, payload)
}
