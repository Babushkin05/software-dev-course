package services

import (
	"context"
	"io"

	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/domain/ports/input"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/domain/ports/output"
)

type FileService struct {
	FileStoringClient  output.FileStoringClient
	FileAnalysisClient output.FileAnalysisClient
}

var _ input.FileUsecase = (*FileService)(nil) // Compile-time check

func NewFileService(fs output.FileStoringClient, fa output.FileAnalysisClient) *FileService {
	return &FileService{
		FileStoringClient:  fs,
		FileAnalysisClient: fa,
	}
}

func (s *FileService) UploadFile(ctx context.Context, file io.Reader, filename string) (string, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return s.FileStoringClient.Upload(ctx, content, filename)
}

func (s *FileService) AnalyzeFile(ctx context.Context, fileID string) ([]byte, string, error) {
	return s.FileAnalysisClient.AnalyzeFile(ctx, fileID)
}

func (s *FileService) DownloadFile(ctx context.Context, fileID string) ([]byte, string, error) {
	return s.FileStoringClient.Download(ctx, fileID)
}

func (s *FileService) GetWordCloud(ctx context.Context, fileID string) ([]byte, string, error) {
	return s.FileAnalysisClient.GetWordCloud(ctx, fileID)
}
