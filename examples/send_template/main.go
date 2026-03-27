package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages"
	"github.com/diegoyosiura/go-whatsapp-cloud/messages/domain"
)

// To run this example:
// 1. Export the variables or use the root .env loader.
// 2. cd examples/send_template
// 3. go run main.go
func main() {
	// 1. Environment Parsing
	version := os.Getenv("VERSION")
	phoneID := os.Getenv("PHONE_NUMBER_ID")
	token := os.Getenv("USER_ACCESS_TOKEN")
	recipient := os.Getenv("RECIPIENT_PHONE_NUMBER")

	if token == "" || recipient == "" {
		log.Fatal("Please populate your env variables before running this example.")
	}

	ctx := context.Background()

	// 2. Initialize the Sending Component
	msgClient := messages.NewClient(version, phoneID, token)

	log.Println("Dispatching Hello World Template to user...")

	// 3. Mount the parameters for the Meta template.
	// If the template possessed variables (Header/Body/Buttons), you would map them in the Components array.
	var components []domain.TemplateComponent

	// 4. Send
	// Note: Signature is (ctx, to, templateName, languageCode, components)
	resp, err := msgClient.SendTemplate(ctx, recipient, "hello_world", "en_US", components)
	if err != nil {
		log.Fatalf("Critical fail during dispatch: %v", err)
	}

	if len(resp.Messages) > 0 {
		fmt.Printf("\nSuccessfully dispatched WhatsApp Template!\nMessage ID tracking token: %s\n", resp.Messages[0].ID)
	}
}
