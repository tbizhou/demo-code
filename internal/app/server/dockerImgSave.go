package server

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

func dockerImgSave(cli *client.Client, ctx context.Context, imageName string) error {
	fmt.Printf("%s starting save\n", imageName)

	saveResponse, err := cli.ImageSave(ctx, []string{imageName})
	if err != nil {
		fmt.Printf("Error saving image: %v\n", err)
		return err
	}

	/*reader := io.TeeReader(saveResponse, os.Stdout)*/

	defer saveResponse.Close()

	fmt.Printf("Starting upload %s\n", imageName)

	if err := ImgStreamToMinio(ctx, imageName, saveResponse); err != nil {
		return err
	}

	return nil
}
