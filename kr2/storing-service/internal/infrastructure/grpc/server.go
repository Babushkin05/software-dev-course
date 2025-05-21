package grpc

import (
	"context"

	filestoringpb "github.com/Babushkin05/software-dev-course/kr2/storing-service/api/gen"
	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/application/ports/input"
)

type FileStoringServer struct {
	filestoringpb.UnimplementedFileStoringServiceServer
	usecase input.FileUseCase
}

func NewFileStoringServer(usecase input.FileUseCase) *FileStoringServer {
	return &FileStoringServer{usecase: usecase}
}

func (s *FileStoringServer) UploadFile(ctx context.Context, req *filestoringpb.UploadRequest) (*filestoringpb.UploadResponse, error) {
	id, err := s.usecase.Upload(req.Filename, req.Content)
	if err != nil {
		return nil, err
	}

	return &filestoringpb.UploadResponse{FileId: id}, nil
}

func (s *FileStoringServer) DownloadFile(ctx context.Context, req *filestoringpb.DownloadRequest) (*filestoringpb.DownloadResponse, error) {
	filename, content, err := s.usecase.Download(req.FileId)
	if err != nil {
		return nil, err
	}

	return &filestoringpb.DownloadResponse{
		Filename: filename,
		Content:  content,
	}, nil
}
