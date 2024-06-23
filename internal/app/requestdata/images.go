package requestdata

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Images struct {
	Images string `form:"images" binding:"required"`
}

type DockerPullImages struct {
	Images []string
}

func BindImages(c *gin.Context) *DockerPullImages {
	var Images Images
	var DockerPullImages DockerPullImages
	if errA := c.ShouldBind(&Images); errA != nil {
		c.String(http.StatusUnprocessableEntity, `the body should be formA`)
	}
	requestimg := Images.Images

	DockerPullImages.Images = strings.Split(requestimg, ",")
	return &DockerPullImages
}
