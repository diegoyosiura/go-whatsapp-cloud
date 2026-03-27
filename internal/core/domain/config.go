package domain

// WhatsAppConfig defines the vital settings required to interact with the Meta WhatsApp Cloud API.
// Based on the Postman Environment variables: Version, UserAccessToken, PhoneNumberID and WABAID.
type WhatsAppConfig struct {
	Version         string
	UserAccessToken string
	PhoneNumberID   string
	WABAID          string
}
