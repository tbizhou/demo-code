package server

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"operator-dev/docker-image-download/internal/config"
	"strings"
)

func ImgStreamToMinio(ctx context.Context, imageName string, reader io.Reader) error {

	minioClient := config.MinioClient()
	objimg := strings.Replace(imageName, ":", "_", -1)
	imgtar := fmt.Sprintf(objimg + ".tar")
	uploadInfo, err := minioClient.PutObject(ctx, "img", imgtar, reader, -1, minio.PutObjectOptions{
		ContentType: "application/x-tar",
	})

	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", uploadInfo.Location, uploadInfo.Size)
	return nil
}
