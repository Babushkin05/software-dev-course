package services

import (
	"time"

	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/application/ports/input"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/application/ports/output"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/domain/entities"
	"github.com/google/uuid"
)

type FileService struct {
	storage  output.StoragePort
	fileRepo output.FileRepo
}

func NewFileService(storage output.StoragePort, fileRepo output.FileRepo) input.FileUseCase {
	return &FileService{
		storage:  storage,
		fileRepo: fileRepo,
	}
}

func (s *FileService) Upload(filename string, content []byte) (string, error) {
	fileID := uuid.New().String()

	file := &entities.File{
		ID:        fileID,
		Filename:  filename,
		Size:      int64(len(content)),
		CreatedAt: time.Now(),
	}

	// Сохраняем файл в Object Storage
	if err := s.storage.SaveFile(fileID, content); err != nil {
		return "", err
	}

	// Сохраняем метаданные в БД
	if err := s.fileRepo.SaveMetadata(file); err != nil {
		return "", err
	}

	return fileID, nil
}

func (s *FileService) Download(fileID string) (string, []byte, error) {
	// Получаем метаданные
	file, err := s.fileRepo.GetMetadata(fileID)
	if err != nil {
		return "", nil, err
	}

	// Получаем файл из Object Storage
	content, err := s.storage.GetFile(fileID)
	if err != nil {
		return "", nil, err
	}

	return file.Filename, content, nil
}
