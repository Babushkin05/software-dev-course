package service_grpc

import (
	"context"

	fileanalysispb "github.com/Babushkin05/software-dev-course/kr2/api-gateway/api/gen/fileanalysis"
	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/domain/ports/output"
)

type FileAnalysisGRPCClient struct {
	client fileanalysispb.FileAnalysisServiceClient
}

func NewFileAnalysisClient(client fileanalysispb.FileAnalysisServiceClient) output.FileAnalysisClient {
	return &FileAnalysisGRPCClient{client: client}
}

func (f *FileAnalysisGRPCClient) AnalyzeFile(ctx context.Context, fileID string) ([]byte, string, error) {
	req := &fileanalysispb.AnalyzeRequest{FileId: fileID}

	resp, err := f.client.Analyze(ctx, req)
	if err != nil {
		return nil, "", err
	}

	return resp.GetContent(), resp.GetFilename(), nil
}

func (f *FileAnalysisGRPCClient) GetWordCloud(ctx context.Context, fileID string) ([]byte, string, error) {
	req := &fileanalysispb.WordCloudRequest{FileId: fileID}

	resp, err := f.client.GetWordCloud(ctx, req)
	if err != nil {
		return nil, "", err
	}

	return resp.GetContent(), resp.GetFilename(), nil
}
