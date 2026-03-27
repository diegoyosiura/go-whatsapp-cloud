package domain

import (
	"testing"
)

func TestWhatsAppConfig_Instantiate(t *testing.T) {
	config := WhatsAppConfig{
		Version:         "v18.0",
		UserAccessToken: "EAABtesttoken123",
		PhoneNumberID:   "123456789",
		WABAID:          "987654321",
	}

	if config.Version != "v18.0" {
		t.Errorf("expected Version v18.0, got %s", config.Version)
	}
	if config.UserAccessToken != "EAABtesttoken123" {
		t.Errorf("expected UserAccessToken EAABtesttoken123, got %s", config.UserAccessToken)
	}
	if config.PhoneNumberID != "123456789" {
		t.Errorf("expected PhoneNumberID 123456789, got %s", config.PhoneNumberID)
	}
	if config.WABAID != "987654321" {
		t.Errorf("expected WABAID 987654321, got %s", config.WABAID)
	}
}
