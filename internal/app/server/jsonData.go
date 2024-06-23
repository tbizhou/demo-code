package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"operator-dev/docker-image-download/pkg/utils"
	"time"
)

func JsonData(c *gin.Context) []string {
	body := c.Request.Body
	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	var imgs []string
	err = json.Unmarshal(data, &imgs)
	if err != nil {
		utils.Response(c, http.StatusUnprocessableEntity, 400, nil, "解析错误")
		return nil
	}
	fmt.Printf("即将拉取以下镜像:\n")
	for _, img := range imgs {
		fmt.Println("name:", img)
	}
	time.Sleep(3 * time.Second)
	return imgs
}
