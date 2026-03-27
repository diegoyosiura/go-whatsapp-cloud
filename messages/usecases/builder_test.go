package usecases

import (
	"testing"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
)

func TestBuildTextPayload(t *testing.T) {
	t.Run("Create generic text message", func(t *testing.T) {
		to := "5511999999999"
		body := "Hello World"

		payload := BuildTextPayload(to, body)

		if payload.MessagingProduct != "whatsapp" {
			t.Errorf("expected messaging_product 'whatsapp', got '%s'", payload.MessagingProduct)
		}
		if payload.To != to {
			t.Errorf("expected to '%s', got '%s'", to, payload.To)
		}
		if payload.Type != "text" {
			t.Errorf("expected type 'text', got '%s'", payload.Type)
		}
		if payload.Text == nil {
			t.Fatalf("expected text object not to be nil")
		}
		if payload.Text.Body != body {
			t.Errorf("expected text body '%s', got '%s'", body, payload.Text.Body)
		}
	})

	t.Run("Create generic template message", func(t *testing.T) {
		to := "123"
		templateName := "hello_world"
		lang := "en_US"

		payload := BuildTemplatePayload(to, templateName, lang, nil)

		if payload.Type != "template" {
			t.Errorf("expected type 'template'")
		}
		if payload.Template == nil || payload.Template.Name != templateName {
			t.Errorf("expected template name %s", templateName)
		}
		if payload.Template.Language.Code != lang {
			t.Errorf("expected lang %s", lang)
		}
	})

	t.Run("Create image message via URL", func(t *testing.T) {
		to := "123"
		imageURL := "https://example.com/img.png"

		payload := BuildImagePayload(to, imageURL, false)

		if payload.Type != "image" {
			t.Errorf("expected type 'image'")
		}
		if payload.Image == nil || payload.Image.Link != imageURL {
			t.Errorf("expected image link %s", imageURL)
		}
	})

	t.Run("Create interactive button message", func(t *testing.T) {
		to := "123"
		text := "Confirm?"
		buttons := []domain.Button{
			{Type: "reply", Reply: domain.Reply{ID: "btn1", Title: "Yes"}},
		}

		payload := BuildInteractiveButtonPayload(to, text, buttons)

		if payload.Type != "interactive" {
			t.Errorf("expected type 'interactive'")
		}
		if payload.Interactive == nil || payload.Interactive.Type != "button" {
			t.Errorf("expected interactive button")
		}
		if payload.Interactive.Body.Text != text {
			t.Errorf("expected interactive body text %s", text)
		}
		if len(payload.Interactive.Action.Buttons) != 1 {
			t.Errorf("expected 1 button")
		}
	})
}
