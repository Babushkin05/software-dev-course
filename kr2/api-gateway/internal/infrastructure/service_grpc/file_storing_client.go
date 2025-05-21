package service_grpc

import (
	"context"

	filestoringpb "github.com/Babushkin05/software-dev-course/kr2/api-gateway/api/gen/filestoring"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/domain/ports/output"
)

type FileStoringGRPCClient struct {
	client filestoringpb.FileStoringServiceClient
}

func NewFileStoringClient(client filestoringpb.FileStoringServiceClient) output.FileStoringClient {
	return &FileStoringGRPCClient{client: client}
}

func (f *FileStoringGRPCClient) Upload(ctx context.Context, content []byte, filename string) (string, error) {
	req := &filestoringpb.UploadRequest{
		Filename: filename,
		Content:  content,
	}

	resp, err := f.client.Upload(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetFileId(), nil
}

func (f *FileStoringGRPCClient) Download(ctx context.Context, fileID string) ([]byte, string, error) {
	req := &filestoringpb.DownloadRequest{
		FileId: fileID,
	}

	resp, err := f.client.Download(ctx, req)
	if err != nil {
		return nil, "", err
	}

	return resp.GetContent(), resp.GetFilename(), nil
}
