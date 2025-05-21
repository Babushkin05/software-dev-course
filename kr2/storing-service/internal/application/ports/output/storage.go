package output

// StoragePort — интерфейс для взаимодействия с yandex object storage
type StoragePort interface {
	SaveFile(fileID string, content []byte) error
	GetFile(fileID string) ([]byte, error)
}
