# Go WhatsApp Cloud API SDK

![Go Version](https://img.shields.io/badge/go-1.18%2B-blue.svg)
![WhatsApp API](https://img.shields.io/badge/WhatsApp%20Cloud-v20.0-25d366.svg)

Um SDK oficial do Golang para interagir com naturalidade com a [API Cloud do WhatsApp Institucional (Business)](https://developers.facebook.com/docs/whatsapp/cloud-api/). 
Projetado focado primariamente em DX (Experiência do Desenvolvedor), usando arquitetura limpa hexagonal, dependências zero de terceiros (`standard library` apenas) e forte resiliência para concorrências e multi-têntencias.

## Visão Geral

O gigantesco escopo de comunicação do WhatsApp Business Manager foi domado em 5 módulos lógicos principais:
1. `messages`: Para o envio síncrono proativo pro usuário (Textos, Mídias, Templates com Botões);
2. `webhook`: Um superservidor/Listener que gerencia a dolorosa etapa de Handshake HTTP GET (Cloud API auth) e o parser contínuo e roteamento de Webhooks JSON (Eventos de leitura, entrega, e resposta do cliente);
3. `media`: Soluções Out-Of-The-Box em *Multipart Streaming* e *Download Encriptado* para hospedar e extrair Voice Notes/PDFs da nuvem da Meta de volta pra você, poupando a sua memória RAM;
4. `businessprofiles` & `waba`: Alterações administrativas sistêmicas da Entidade da Conta;
5. `phonenumbers`: Trinca o Setup do P.I.N para ativação de um Número sob o escopo WABA da corporação.

---

## 🚀 Como Iniciar

### 1. Instalação
Incorpore-o em seu `go.mod` facilmente:
```bash
go get -u github.com/diegoyosiura/go-whatsapp-cloud
```

### 2. Disparando seu Primeiro Template MSG
Veja como a complexidade inteira do payload e tokens se restringe a poucas linhas:

```go
package main

import (
	"context"
	"os"

	"github.com/diegoyosiura/go-whatsapp-cloud/messages"
)

func main() {
    // Inject API Keys
	client := messages.NewClient("v20.0", os.Getenv("PHONE_ID"), os.Getenv("USER_TOKEN"))

	ctx := context.Background()
	_, err := client.SendTemplate(ctx, "+551199999999", "hello_world", "en_US", nil)
	if err != nil {
		panic(err)
	}
}
```

### 3. Subindo o seu próprio Bot (Webhook Receiver)
Receba e parseie eventos em apenas 1 linha de `http.ServeMux`! Todos os Eventos como leituras (`READ`), mensagens recebidas, Status e mídias vêm convertidos para Go Structs limpos, checados criptograficamente.

```go
import "github.com/diegoyosiura/go-whatsapp-cloud/webhook"

func main() {
	client := webhook.NewClient(os.Getenv("SECRET"), os.Getenv("VERIFY_TOKEN"))

	mux := http.NewServeMux()
	mux.Handle("/webhook", client.HTTPHandler())

	http.ListenAndServe(":8080", mux)
}
```

## Explorando Mais
Acessem a pasta absoluta de base (`/examples`) no código fonte do SDK! Encontrem:
- **`examples/send_template/main.go`**: Fluxo extenso de comunicação com o Usuário Final;
- **`examples/webhook_bot/main.go`**: Mapeando e logando JSON payloads no terminal.

### Arquitetura Customizável
O SDK utiliza Portas e Adaptadores (Hexagonal). Você não é obrigado a usar os módulos padrões incluídos na pasta `internal/adapters`. Por exemplo, a injeção nativa de Storage salva fotos no disco em C:/ ou /tmp/, mas se você desejar hospedar as mídias em um bucket S3 da AWS, basta criar seu próprio client do S3 amarrado à abstração visual do pacote *Ports* do Go Whatsapp SDK.
