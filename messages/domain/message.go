package domain

// SendMessagePayload is the main struct strictly reflecting the WhatsApp Graph API JSON structure.
type SendMessagePayload struct {
	MessagingProduct string          `json:"messaging_product"`
	RecipientType    string          `json:"recipient_type,omitempty"`
	To               string          `json:"to"`
	Type             string          `json:"type"`
	Text             *TextObject     `json:"text,omitempty"`
	Template         *TemplateObject `json:"template,omitempty"`
	Image            *MediaObject    `json:"image,omitempty"`
	Video            *MediaObject    `json:"video,omitempty"`
	Document         *MediaObject    `json:"document,omitempty"`
	Interactive      *InteractiveObj `json:"interactive,omitempty"`
}

type MediaObject struct {
	ID   string `json:"id,omitempty"`
	Link string `json:"link,omitempty"`
}

type InteractiveObj struct {
	Type   string          `json:"type"`
	Body   InteractiveBody `json:"body"`
	Action Action          `json:"action"`
}

type InteractiveBody struct {
	Text string `json:"text"`
}

type Action struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Type  string `json:"type"`
	Reply Reply  `json:"reply"`
}

type Reply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type TextObject struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type TemplateObject struct {
	Name       string              `json:"name"`
	Language   TemplateLanguage    `json:"language"`
	Components []TemplateComponent `json:"components,omitempty"`
}

type TemplateComponent struct {
	Type       string             `json:"type"`
	Parameters []TemplateParameter `json:"parameters"`
}

type TemplateParameter struct {
	Type  string       `json:"type"`
	Text  string       `json:"text,omitempty"`
	Image *MediaObject `json:"image,omitempty"`
}

type TemplateLanguage struct {
	Code string `json:"code"`
}

// MessageResponse represents the Meta's response acknowledging the message transmission.
type MessageResponse struct {
	MessagingProduct string            `json:"messaging_product"`
	Contacts         []ResponseContact `json:"contacts"`
	Messages         []ResponseMessage `json:"messages"`
}

type ResponseContact struct {
	Input string `json:"input"`
	WaID  string `json:"wa_id"`
}

type ResponseMessage struct {
	ID string `json:"id"`
}
