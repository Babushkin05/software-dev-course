package output

import "context"

// StoragePort — интерфейс для взаимодействия с yandex object storage
type StoragePort interface {
	SaveFile(ctx context.Context, fileID string, content []byte) error
	GetFile(ctx context.Context, fileID string) ([]byte, error)
}
