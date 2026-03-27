package ports

import "context"

// FileStorage define o contrato para manipulação de arquivos abstratos (disco, S3, etc).
type FileStorage interface {
	Save(ctx context.Context, path string, data []byte) error
	Get(ctx context.Context, path string) ([]byte, error)
	Delete(ctx context.Context, path string) error
}
