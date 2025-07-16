package output

import (
	"context"
)

type FileStoringClient interface {
	Upload(ctx context.Context, content []byte, filename string) (string, error)
	Download(ctx context.Context, fileID string) ([]byte, string, error)
}
