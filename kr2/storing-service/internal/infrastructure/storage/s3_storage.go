package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Storage struct {
	client *minio.Client
	bucket string
}

func NewS3Storage(c config.S3Config) (*S3Storage, error) {
	endpoint := c.Endpoint
	accessKey := c.Endpoint
	secretKey := c.SecretKey
	useSSL := c.UseSSL
	bucket := c.Bucket

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
		Region: c.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize minio client: %w", err)
	}

	return &S3Storage{client: client, bucket: bucket}, nil
}

func (s *S3Storage) SaveFile(ctx context.Context, id string, data []byte) error {
	reader := bytes.NewReader(data)

	_, err := s.client.PutObject(ctx, s.bucket, id, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: "text/plain",
	})
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	log.Printf("Successfully uploaded file with ID: %s", id)
	return nil
}

func (s *S3Storage) GetFile(ctx context.Context, id string) ([]byte, error) {
	obj, err := s.client.GetObject(ctx, s.bucket, id, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	defer obj.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, obj); err != nil {
		return nil, fmt.Errorf("failed to read object: %w", err)
	}

	return buf.Bytes(), nil
}
