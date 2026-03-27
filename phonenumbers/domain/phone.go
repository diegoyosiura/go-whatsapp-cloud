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

// SetTwoStepVerificationPayload maps the pin definition request.
type SetTwoStepVerificationPayload struct {
	Pin string `json:"pin"`
}

// BlockUserRequest wraps the list of users you want to block or unblock.
type BlockUserRequest struct {
	MessagingProduct string            `json:"messaging_product"`
	BlockUsers       []BlockUserDetail `json:"block_users"`
}

// BlockUserDetail represents the nested user string inside the block flow.
type BlockUserDetail struct {
	User string `json:"user"`
}

// GetBlockedUsersResponse fetches the currently blocked numbers.
type GetBlockedUsersResponse struct {
	Data []GetBlockedUsersData `json:"data"`
}

// GetBlockedUsersData holds individual blocked number properties.
type GetBlockedUsersData struct {
	MessagingProduct string `json:"messaging_product"`
	WAID             string `json:"wa_id"`
}
