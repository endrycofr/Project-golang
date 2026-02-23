package minio

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"bwastartup/config" // sesuaikan path project kamu

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Connect membuat koneksi MinIO dan memastikan bucket tersedia
func Connect(cfg *config.ConfigMinio) (*MinioStorage, error) {

	// Transport untuk handle self-signed TLS dari Kubernetes NodePort
	transport := &http.Transport{}

	if cfg.Secure {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure:    cfg.Secure,
		Transport: transport,
	})

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	exists, err := minioClient.BucketExists(ctx, cfg.BucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, cfg.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
		log.Printf("✅ Bucket created: %s\n", cfg.BucketName)
	}
	log.Println("✅ MinIO Connected")

	return &MinioStorage{
		Client:     minioClient,
		BucketName: cfg.BucketName,
	}, nil
}
