package adapters

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/diegoyosiura/go-whatsapp-cloud/internal/core/ports"
)

type envConfigReader struct{}

// NewEnvConfigReader cria um adaptador de leitura baseado em variáveis de ambiente.
// Retorna a interface ports.ConfigReader para manter a restrição de visibilidade.
func NewEnvConfigReader() ports.ConfigReader {
	return &envConfigReader{}
}

func (e *envConfigReader) GetString(ctx context.Context, key string) (string, error) {
	val := os.Getenv(key)
	return val, nil
}

func (e *envConfigReader) GetInt(ctx context.Context, key string) (int, error) {
	val := os.Getenv(key)
	if val == "" {
		return 0, fmt.Errorf("configuração não encontrada: %s", key)
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("falha ao converter para int: %w", err)
	}

	return i, nil
}

func (e *envConfigReader) GetBool(ctx context.Context, key string) (bool, error) {
	val := os.Getenv(key)
	if val == "" {
		return false, fmt.Errorf("configuração não encontrada: %s", key)
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, fmt.Errorf("falha ao converter para bool: %w", err)
	}

	return b, nil
}
