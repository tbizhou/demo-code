package server

import (
	"context"
	"fmt"
	"github.com/demo-code/internal/config"
	"github.com/minio/minio-go/v7"
	"io"
)

func ImgStreamToMinio(ctx context.Context, imageName string, reader io.Reader) error {

	minioClient := config.MinioClient()
	imgtar := fmt.Sprintf(imageName + ".tar")
	uploadInfo, err := minioClient.PutObject(ctx, "img", imgtar, reader, -1, minio.PutObjectOptions{
		ContentType: "application/x-tar",
	})

	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", uploadInfo.ETag, uploadInfo.Size)
	return nil
}
