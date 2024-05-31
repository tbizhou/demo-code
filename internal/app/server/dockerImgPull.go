package server

import (
	"context"
	"fmt"
	"github.com/demo-code/internal/utils"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sync"
)

var Images []string

func RunPull(ctx *gin.Context) {
	if err := dockerImgPull(); err != nil {
		utils.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "上传错误")
		return
	}
}

func dockerImgPull() error {
	ctx := context.Background()
	Images = []string{
		"alpine:latest",
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, img := range Images {
		wg.Add(1)
		go func() error {
			defer wg.Done()
			pullResponse, _ := cli.ImagePull(ctx, img, image.PullOptions{})

			defer pullResponse.Close()

			if _, err := io.Copy(os.Stdout, pullResponse); err != nil {
				return err
			}

			fmt.Printf("%s pull success\n", img)
			if err := dockerImgSave(cli, ctx, img); err != nil {
				return err
			}
			return nil
		}()
		wg.Wait()
	}
	return nil
	// 等待所有协程完成
}
