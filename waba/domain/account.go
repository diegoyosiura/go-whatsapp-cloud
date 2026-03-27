package domain

// AccountInfo represents metadata fetched from the WhatsApp Business Account global edge.
type AccountInfo struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	Currency                 string `json:"currency"`
	TimezoneID               string `json:"timezone_id"`
	MessageTemplateNamespace string `json:"message_template_namespace"`
}
