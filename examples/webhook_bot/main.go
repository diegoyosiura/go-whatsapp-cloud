package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/diegoyosiura/go-whatsapp-cloud/webhook"
)

// To run this example:
// 1. Export variables APP_SECRET and WEBHOOK_VERIFY_TOKEN
// 2. cd examples/webhook_bot
// 3. go run main.go
func main() {
	appSecret := os.Getenv("APP_SECRET")
	verifyToken := os.Getenv("WEBHOOK_VERIFY_TOKEN") // Used only for initial GET handshake

	if appSecret == "" {
		log.Println("Warning: APP_SECRET is empty. Signature security is degraded.")
	}

	// 1. Initialize the Webhook Client. The webhook dispatcher routes payloads securely.
	client := webhook.NewClient(appSecret, verifyToken)

	// 2. Setting up Route logic
	mux := http.NewServeMux()

	// 2a. Register the Go standard HTTP Handler seamlessly exposed by the SDK Facade.
	// This single line automatically handles GET (Challenge Handshake) and POST (Event Processing & Signature Checking).
	mux.Handle("/webhook", client.HTTPHandler())

	fmt.Println("Starting Webhook Server at port 8080...")
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
