package output

import "context"

type StoringClient interface {
	DownloadFile(ctx context.Context, fileID string) (filename string, content []byte, err error)
}
