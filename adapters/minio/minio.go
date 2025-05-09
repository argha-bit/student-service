package minio

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func GetMinIOClient() (*minio.Client, error) {
	var err error
	server := os.Getenv("MINIO_SERVER")

	log.Println("MinIO is connecting on", server)
	minioClient, err := minio.New(server, &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_USERNAME"), os.Getenv("MINIO_PASSWORD"), ""),
		Secure: false, // Set to true for HTTPS
		Region: os.Getenv("MINIO_REGION"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to MinIO: %v", err)
		return minioClient, err
	} else {
		log.Println("MinIO connected successfully")
	}
	return minioClient, err
}
