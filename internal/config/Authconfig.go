package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/registry"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"operator-dev/docker-image-download/internal/modules"
)

var c, _ = modules.ReadConfig()

func EncodedJson() string {
	DockerConfig := c.Docker
	authConfig := registry.AuthConfig{
		Username: DockerConfig.Username,
		Password: DockerConfig.Password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(encodedJSON)
}

func MinioClient() *minio.Client {
	MinioConfig := c.Minio
	minioEndpoint := MinioConfig.Endpoint
	accessKeyID := MinioConfig.AccessKey
	secretAccessKey := MinioConfig.SecretKey
	useSSL := MinioConfig.SSL
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
