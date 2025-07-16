package output

import "context"

type WordCloudGenerator interface {
	GenerateImage(ctx context.Context, text string) ([]byte, error)
}
