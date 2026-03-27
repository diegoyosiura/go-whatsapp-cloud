package usecases

import "github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"

// BuildTextPayload constructs a valid Text Message payload for the WhatsApp API.
func BuildTextPayload(to string, body string) domain.SendMessagePayload {
	return domain.SendMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "text",
		Text: &domain.TextObject{
			PreviewURL: false,
			Body:       body,
		},
	}
}

// BuildTemplatePayload creates a payload for Meta approved templates.
func BuildTemplatePayload(to, templateName, languageCode string, components []domain.TemplateComponent) domain.SendMessagePayload {
	return domain.SendMessagePayload{
		MessagingProduct: "whatsapp",
		To:               to,
		Type:             "template",
		Template: &domain.TemplateObject{
			Name: templateName,
			Language: domain.TemplateLanguage{
				Code: languageCode,
			},
			Components: components,
		},
	}
}

// BuildImagePayload creates a payload for sending an image from an external URL.
func BuildImagePayload(to, imageURL string, isID bool) domain.SendMessagePayload {
	img := &domain.MediaObject{}
	if isID {
		img.ID = imageURL
	} else {
		img.Link = imageURL
	}

	return domain.SendMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "image",
		Image:            img,
	}
}

// BuildInteractiveButtonPayload creates a payload for sending quick reply buttons.
func BuildInteractiveButtonPayload(to string, bodyText string, buttons []domain.Button) domain.SendMessagePayload {
	return domain.SendMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "interactive",
		Interactive: &domain.InteractiveObj{
			Type: "button",
			Body: domain.InteractiveBody{
				Text: bodyText,
			},
			Action: domain.Action{
				Buttons: buttons,
			},
		},
	}
}
