package domain

// RequestCodePayload maps the requirement for the Meta Graph endpoint used to request an onboarding setup PIN.
type RequestCodePayload struct {
	CodeMethod string `json:"code_method"` // SMS or VOICE
	Language   string `json:"language"`    // e.g., en, pt_BR
}

// VerifyCodePayload mirrors the JSON payload sent to the graph immediately after receiving the user's setup PIN.
type VerifyCodePayload struct {
	Code string `json:"code"`
}
