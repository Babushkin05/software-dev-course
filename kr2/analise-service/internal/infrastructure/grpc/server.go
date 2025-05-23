package grpc

import (
	"context"
	"log"

	fileanalysispb "github.com/Babushkin05/software-dev-course/kr2/analise-service/api/gen/fileanalisys"
	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/ports/input"
)

type analysisGrpcServer struct {
	fileanalysispb.UnimplementedFileAnalysisServiceServer
	service input.AnalysisUseCase
}

func NewGRPCServer(service input.AnalysisUseCase) fileanalysispb.FileAnalysisServiceServer {
	return &analysisGrpcServer{service: service}
}

func (s *analysisGrpcServer) Analyze(ctx context.Context, req *fileanalysispb.AnalyzeRequest) (*fileanalysispb.AnalyzeResponse, error) {
	filename, content, err := s.service.Analyze(ctx, req.GetFileId())
	if err != nil {
		log.Printf("Analyze error: %v", err)
		return nil, err
	}

	return &fileanalysispb.AnalyzeResponse{
		Filename: filename,
		Content:  []byte(content),
	}, nil
}

func (s *analysisGrpcServer) GetWordCloud(ctx context.Context, req *fileanalysispb.WordCloudRequest) (*fileanalysispb.WordCloudResponse, error) {
	filename, imgData, err := s.service.GenerateWordCloud(ctx, req.GetFileId())
	if err != nil {
		log.Printf("WordCloud error: %v", err)
		return nil, err
	}

	return &fileanalysispb.WordCloudResponse{
		Filename: filename,
		Content:  imgData,
	}, nil
}
