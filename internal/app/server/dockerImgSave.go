package server

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"io"
	"os"
)

func dockerImgSave(cli *client.Client, ctx context.Context, imageName string) error {

	saveResponse, err := cli.ImageSave(ctx, []string{imageName})
	if err != nil {
		fmt.Printf("Error saving image: %v\n", err)
		return err
	}
	fmt.Println("Starting images stream")
	reader := io.TeeReader(saveResponse, os.Stdout)

	defer saveResponse.Close()

	fmt.Println("Starting upload")

	if err := ImgStreamToMinio(ctx, imageName, reader); err != nil {
		return err
	}

	return nil
}
