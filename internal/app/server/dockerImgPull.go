package server

import (
	"context"
	"fmt"
	"github.com/demo-code/internal/app/requestdata"
	"github.com/demo-code/internal/utils"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func RunPull(c *gin.Context) {
	images := requestdata.BindImages(c)
	if err := dockerImgPull(images); err != nil {
		utils.Response(c, http.StatusUnprocessableEntity, 400, nil, "上传错误")
		return
	}
}

func dockerImgPull(Images *requestdata.DockerPullImages) error {
	ctx := context.Background()

	time.Sleep(5 * time.Second)
	fmt.Println("Images:", Images.Images)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, img := range Images.Images {
		fmt.Println("img:", img)
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
