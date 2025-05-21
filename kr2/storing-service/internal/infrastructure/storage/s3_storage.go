package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Storage struct {
	client *minio.Client
	bucket string
}

func NewS3Storage() (*S3Storage, error) {
	endpoint := os.Getenv("S3_ENDPOINT")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")
	useSSLStr := os.Getenv("S3_USE_SSL")
	bucket := os.Getenv("S3_BUCKET")

	useSSL, err := strconv.ParseBool(useSSLStr)
	if err != nil {
		return nil, fmt.Errorf("invalid S3_USE_SSL value: %v", err)
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
		Region: os.Getenv("S3_REGION"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize minio client: %w", err)
	}

	return &S3Storage{client: client, bucket: bucket}, nil
}

func (s *S3Storage) Save(ctx context.Context, id string, data []byte) error {
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

func (s *S3Storage) Get(ctx context.Context, id string) ([]byte, error) {
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
