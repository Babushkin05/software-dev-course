package output

import (
	"context"
)

type FileAnalysisClient interface {
	Analyze(ctx context.Context, fileID string) ([]byte, string, error)
	GetWordCloud(ctx context.Context, fileID string) ([]byte, string, error)
}
