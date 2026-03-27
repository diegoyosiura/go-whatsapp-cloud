package domain

// QRCode represents the core entity for interacting with physical or deep-link chats.
type QRCode struct {
	Code             string `json:"code"`
	PrefilledMessage string `json:"prefilled_message"`
	DeepLinkURL      string `json:"deep_link_url,omitempty"`
	ImageURL         string `json:"qr_image_url,omitempty"`
}

// QRCodeListResponse maps the collection array when loading multiple codes.
type QRCodeListResponse struct {
	Data []QRCode `json:"data"`
}

// SuccessResponse defines boolean returns for Delete actions.
type SuccessResponse struct {
	Success bool `json:"success"`
}

// CreateQRCodeRequest payload to ask for a new QR.
type CreateQRCodeRequest struct {
	PrefilledMessage string `json:"prefilled_message"`
	GenerateQRImage  string `json:"generate_qr_image"`
}

// UpdateQRCodeRequest updates the message of an existing code.
type UpdateQRCodeRequest struct {
	Code             string `json:"code"`
	PrefilledMessage string `json:"prefilled_message"`
}
