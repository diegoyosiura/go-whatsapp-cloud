package ports

import "context"

// ConfigReader define o contrato para leitura de configurações geradas externamente.
type ConfigReader interface {
	GetString(ctx context.Context, key string) (string, error)
	GetInt(ctx context.Context, key string) (int, error)
	GetBool(ctx context.Context, key string) (bool, error)
}
