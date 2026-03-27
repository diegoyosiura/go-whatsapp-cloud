package ports

import "context"

// PhoneService acts as an intermediator contract abstracting Meta parameters to the OTP engine.
type PhoneService interface {
	RequestCode(ctx context.Context, codeMethod, language string) error
	VerifyCode(ctx context.Context, code string) error
}
