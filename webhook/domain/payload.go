package domain

// WebhookPayload represents the root JSON structure sent by Meta.
type WebhookPayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         Metadata   `json:"metadata"`
	Contacts         []Contact  `json:"contacts,omitempty"`
	Messages         []Message  `json:"messages,omitempty"`
	Statuses         []Status   `json:"statuses,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Text      *Text  `json:"text,omitempty"`
	// Additional types will be mapped later, currently handling core attributes and Text.
}

type Text struct {
	Body string `json:"body"`
}

type Status struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Recipient string `json:"recipient_id"`
}
