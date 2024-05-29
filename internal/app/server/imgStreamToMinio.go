package server

import (
	"context"
	"fmt"
	"github.com/demo-code/internal/config"
	"github.com/minio/minio-go/v7"
	"io"
)

func ImgStreamToMinio(ctx context.Context, imageName string, reader io.Reader) error {
	fmt.Println("connect minio")
	minioClient := config.MinioClient()
	fmt.Sprintf(imageName)
	uploadInfo, err := minioClient.PutObject(ctx, "images", "nginx_1.25.5.tar", reader, -1, minio.PutObjectOptions{
		ContentType: "application/x-tar",
	})

	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", uploadInfo.Key, uploadInfo.Size)
	return nil
}
