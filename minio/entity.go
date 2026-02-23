package minio

import "github.com/minio/minio-go/v7"

type MinioStorage struct {
	Client     *minio.Client
	BucketName string
}
