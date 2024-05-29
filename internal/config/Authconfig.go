package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/registry"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func EncodedJson() string {
	authConfig := registry.AuthConfig{
		Username: "ansible",
		Password: "only_pull_!#@",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(encodedJSON)
}

func MinioClient() *minio.Client {
	minioEndpoint := "10.10.10.12:9000"
	accessKeyID := "sign-minio"
	secretAccessKey := "whoami?Minio@123"
	useSSL := false
	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Printf("Error initializing Minio client: %v\n", err)
		return nil
	}
	return minioClient
}
