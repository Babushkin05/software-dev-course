package input

import (
	"context"
	"io"
)

type FileUsecase interface {
	UploadFile(ctx context.Context, file io.Reader, filename string) (string, error)
	AnalyzeFile(ctx context.Context, fileID string) ([]byte, error)
	DownloadFile(ctx context.Context, fileID string) ([]byte, string, error)
	GetWordCloud(ctx context.Context, fileID string) ([]byte, string, error)
}
