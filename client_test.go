package whatsapp

import (
	"testing"

	"github.com/diegoyosiura/go-whatsapp-cloud/internal/core/domain"
)

func TestNewClient(t *testing.T) {
	config := domain.WhatsAppConfig{
		Version:         "v18.0",
		UserAccessToken: "EAABtesttoken",
		PhoneNumberID:   "123",
		WABAID:          "456",
	}

	client := NewClient(config)
	if client == nil {
		t.Fatal("expected client to not be nil")
	}

	// Verify that the config passed in the constructor is accessible
	c, err := client.GetPhoneNumberConfig("123")
	if err != nil {
		t.Fatalf("expected no error getting primary config, got %v", err)
	}
	if c.PhoneNumberID != "123" {
		t.Errorf("expected PhoneNumberID to be 123, got %s", c.PhoneNumberID)
	}
}

func TestClient_AddPhoneNumberConfig(t *testing.T) {
	config1 := domain.WhatsAppConfig{
		Version:         "v18.0",
		UserAccessToken: "token1",
		PhoneNumberID:   "111",
		WABAID:          "WABA1",
	}
	client := NewClient(config1)

	config2 := domain.WhatsAppConfig{
		Version:         "v18.0",
		UserAccessToken: "token2",
		PhoneNumberID:   "222",
		WABAID:          "WABA1",
	}
	
	// Add secondary config
	client.AddPhoneNumberConfig("222", config2)

	// Retrieve secondary config
	c, err := client.GetPhoneNumberConfig("222")
	if err != nil {
		t.Fatalf("expected to find config for 222, got error: %v", err)
	}
	if c.UserAccessToken != "token2" {
		t.Errorf("expected token2, got %s", c.UserAccessToken)
	}

	// Retrieve non-existent config should fail
	_, err = client.GetPhoneNumberConfig("999")
	if err == nil {
		t.Error("expected error when getting non-existent config, got nil")
	}
}
