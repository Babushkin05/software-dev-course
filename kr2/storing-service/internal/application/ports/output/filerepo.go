package output

import "github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/domain/entities"

type FileRepo interface {
	SaveMetadata(file *entities.File) error
	GetMetadata(fileID string) (*entities.File, error)
}
