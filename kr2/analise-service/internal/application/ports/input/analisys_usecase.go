package input

import "context"

type AnalysisUseCase interface {
	Analyze(ctx context.Context, fileID string) (filename string, summary string, err error)

	GenerateWordCloud(ctx context.Context, fileID string) (filename string, image []byte, err error)
}
