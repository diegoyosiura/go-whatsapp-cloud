# Go WhatsApp Cloud API SDK

![Go Version](https://img.shields.io/badge/go-1.18%2B-blue.svg)
![WhatsApp API](https://img.shields.io/badge/WhatsApp%20Cloud-v20.0-25d366.svg)
![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen.svg)

Um SDK oficial em Golang para interagir fluida e estaticamente com a [API Cloud do WhatsApp Business](https://developers.facebook.com/docs/whatsapp/cloud-api/).
Construído com foco resoluto em **DX (Developer Experience)**, este SDK utiliza uma arquitetura limpa (Hexagonal/VIPER) estrita e garante dependência ZERO de frameworks externos (unicamente a **Standard Library** do Go), imune a quebras de compatibilidade.

---

## 🚀 Como Iniciar

### 1. Instalação
Incorpore-o ao seu módulo executando:
```bash
go get -u github.com/diegoyosiura/go-whatsapp-cloud
```

### 2. Configurando o Client Master (Multi-Tenant)
A arquitetura foi refinada para suporte nativo assíncrono a múltiplos números de telefone (`tenant`) de forma transparente sob o mesmo servidor.

```go
package main

import (
	"context"
	"os"

	whatsapp "github.com/diegoyosiura/go-whatsapp-cloud"
	"github.com/diegoyosiura/go-whatsapp-cloud/internal/core/domain"
)

func main() {
	// Inicialize a Configuração Primária do Hub
	client := whatsapp.NewClient(domain.WhatsAppConfig{
		Version:         "v20.0",
		UserAccessToken: os.Getenv("USER_TOKEN"),
		PhoneNumberID:   os.Getenv("PHONE_ID"),
		WABAID:          os.Getenv("WABA_ID"), // Exclusivo para métricas WABA/QR
	})

	// (Opcional) Adicione suporte a um segundo número na mesma instância!
	client.AddPhoneNumberConfig("PHONE_ID_2", domain.WhatsAppConfig{
        Version: "v20.0",
        UserAccessToken: "TOKEN_2",
        PhoneNumberID: "PHONE_ID_2",
    })
```

### 3. Disparando Mensagens (Fluent API)
O SDK expõe métodos encadeados lógicos e blindados contextualmente (Fluent Builders). Para mandar Textos via Instância Primária:

```go
	ctx := context.Background()
	_, err := client.Primary().Messages().SendText(ctx, "+551199999999", "Olá do Go SDK!")
	if err != nil {
		panic(err)
	}
    
    // Trabalhando com a Instância Secundária do mesmo Client Global
    client.ForPhone("PHONE_ID_2").Messages().SendText(ctx, "+551188888888", "Mensagem da filial!")
}
```

### 4. Subindo o seu próprio Bot (Webhook Receiver)
Receba e escute eventos da API Cloud em apenas 1 linha de `http.ServeMux`! Webhooks chegam desestruturados em Structs Go perfeitamente tipadas após cruzamento criptográfico `X-Hub-Signature`.

```go
	// Webhooks são agnósticos de telefone e configurados à nível de App (Meta Dev)
	whClient := client.Webhook(os.Getenv("APP_SECRET"), os.Getenv("VERIFY_TOKEN"))

	mux := http.NewServeMux()
	mux.Handle("/webhook", whClient.HTTPHandler())

	http.ListenAndServe(":8080", mux)
```

---

## 📦 Módulos Suportados (API Coverage)
Coberto por Table-Driven TDD's absolutos, o Core SDK abrange o espectro massivo do Graph:

- **`Messages`**: Textos, Mídias e Templates Dinâmicos (*Interactive Actions*).
- **`Media`**: Upload e Download Seguro em buffer rotativo restrito.
- **`Webhook`**: Listener robusto de Statuses com Validação de Identidade Criptografada.
- **`QRCodes`**: Automação profunda via prefilled messages de QR Codes de Lead.
- **`Analytics`**: Extratores de métricas (*Billing Throughput*) focadas em Conta de Negócio ou Conversa.
- **`PhoneNumbers`**: Bloqueios automatizados Anti-Spam (Block/Unblock) e Setups de Two-Step PIN (2FA).
- **`BusinessProfiles`**: Gestão do Perfil Meta Business Público programaticamente (About, Addresses).
- **`Uploads` (Resumable)**: Envio via Chunk sessions escalável para arquivos maciços corporativos (+100MB).

---

## 🏗️ Arquitetura Personalizável (Hexagonal Ports)
Você não está acorrentado aos padrões locais! Se pretende gravar logs ou transbordar os bytes brutos do Whatsapp diretamente num Amazon S3 Buckets ou no Cloud SQL, a estrutura de interfaces nas pastas `/ports` aceita facilmente a injeção do seu Adapter caseiro substituindo a engine base sem quebrar o framework de Mensageria.

Licenciado via MIT.
