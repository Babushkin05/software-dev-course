package output

import (
	"context"
)

type FileAnalysisClient interface {
	AnalyzeFile(ctx context.Context, fileID string) ([]byte, string, error)
	GetWordCloud(ctx context.Context, fileID string) ([]byte, string, error)
}
