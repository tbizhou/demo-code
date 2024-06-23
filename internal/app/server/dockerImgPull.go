package server

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"operator-dev/docker-image-download/internal/config"
	"operator-dev/docker-image-download/pkg/utils"
	"os"
	"sync"
)

func RunPull(c *gin.Context) {
	//images := requestdata.BindImages(c)
	images := JsonData(c)
	if err := dockerImgPull(images); err != nil {
		utils.Response(c, http.StatusUnprocessableEntity, 400, nil, "上传错误")
		return
	}
}

func RunDockerImgPull(Img string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		<-sem
	}()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		return
	}
	pullResponse, err := cli.ImagePull(ctx, Img, image.PullOptions{RegistryAuth: config.EncodedJson()})
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := io.Copy(os.Stdout, pullResponse); err != nil {
		fmt.Println(err)
		return
	}
	defer pullResponse.Close()
	dockerImgSave(cli, ctx, Img)

}

func dockerImgPull(Images []string) error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 3) // 限制并发数为 3
	for _, img := range Images {
		fmt.Println("start image pull:", img)
		wg.Add(1)
		sem <- struct{}{}
		img := img
		go RunDockerImgPull(img, &wg, sem)
	}
	wg.Wait()
	return nil

}
