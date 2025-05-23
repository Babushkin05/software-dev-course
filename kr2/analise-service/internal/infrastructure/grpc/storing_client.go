package grpc

import (
	"context"
	"fmt"

	filestoringpb "github.com/Babushkin05/software-dev-course/kr2/analise-service/api/gen/filestoring"
	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/ports/output"
	"google.golang.org/grpc"
)

type storingClient struct {
	client filestoringpb.FileStoringServiceClient
}

func NewStoringClient(conn *grpc.ClientConn) output.StoringClient {
	return &storingClient{
		client: filestoringpb.NewFileStoringServiceClient(conn),
	}
}

func (s *storingClient) DownloadFile(ctx context.Context, fileID string) (string, []byte, error) {
	resp, err := s.client.Download(ctx, &filestoringpb.DownloadRequest{FileId: fileID})
	if err != nil {
		return "", nil, fmt.Errorf("failed to download file from storing service: %w", err)
	}

	return resp.Filename, resp.Content, nil
}
