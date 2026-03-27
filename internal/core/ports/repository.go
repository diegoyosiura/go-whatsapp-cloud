package ports

import "context"

// BaseRepository define o contrato primário para checagens baseadas em estado ou persistência.
// Em cenários de negócio avançados, essa interface pode ser combinada ou estendida.
type BaseRepository interface {
	Ping(ctx context.Context) error
	SaveLog(ctx context.Context, log interface{}) error
}
