package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages"
)

func main() {
	envFile, err := os.Open("../../.env")
	if err != nil {
		fmt.Printf("Failed to open .env: %v\n", err)
		return
	}
	defer envFile.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			envVars[strings.TrimSpace(parts[0])] = strings.Trim(strings.TrimSpace(parts[1]), `"`)
		}
	}

	apiVersion := envVars["Version"]
	phoneNumberID := envVars["Phone-NumberID"]
	token := envVars["UserAccessToken"]
	recipient := envVars["RecipientPhoneNumber"] // 5511974295577

	client := messages.NewClient(apiVersion, phoneNumberID, token)

	fmt.Println("Sending real message to:", recipient, "via phoneId:", phoneNumberID)
	resp, err := client.SendText(context.Background(), recipient, "Olá Mestre! Aqui é a IA do Antigravity. O mecanismo do motor de envio de mensagens do seu SDK Go (whatsapp-go-sdk) acaba de ser testado com absoluto sucesso de ponta a ponta e acaba de entregar esta notificação. Fase 3 inicializada!")

	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		return
	}

	fmt.Printf("Message Sent successfully! Meta Response: %+v\n", resp)
}
