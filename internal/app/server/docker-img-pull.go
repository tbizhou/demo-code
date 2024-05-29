package server

import (
	"context"
	"fmt"
	"github.com/demo-code/internal/utils"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"

	"os"
	"sync"
)

func RunPull(ctx *gin.Context) {
	if err := StreamDockerImage(); err != nil {
		utils.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "上传错误")
		return
	}
}

func StreamDockerImage() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//docker registry username password json处理
	//authStr := config.EncodedJson()
	var wg sync.WaitGroup
	wg.Add(1)
	//r, w := io.Pipe()
	go func() {
		defer wg.Done()
		//pullOptions := image.PullOptions{RegistryAuth: authStr}
		pullResponse, err := cli.ImagePull(context.Background(), "nginx:1.25.5", image.PullOptions{})
		if err != nil {
			fmt.Printf("Error pulling image: %v\n", err)
			return
		}
		fmt.Println("image pull success")

		defer pullResponse.Close()

		if _, err := io.Copy(os.Stdout, pullResponse); err != nil {
			log.Fatalln(err)
		}
		saveResponse, err := cli.ImageSave(context.Background(), []string{"nginx:1.25.5"})
		if err != nil {
			fmt.Printf("Error saving image: %v\n", err)
			return
		}
		reader := io.TeeReader(saveResponse, os.Stdout)
		defer saveResponse.Close()
		fmt.Println("image save success")

		/*		if _, err := io.Copy(w, saveResponse); err != nil {
				log.Fatalln(err)
			}*/

		fmt.Println("start upload")
		if err := StreamImgToMinio(context.Background(), reader); err != nil {
			return
		}

	}()
	wg.Wait()
	/*	if err := StreamImgToMinio(context.Background(), r); err != nil {
		return err
	}*/
	fmt.Println("pull image is done")
	return nil
}
