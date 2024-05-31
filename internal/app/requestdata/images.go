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

func BindImages(ctx *gin.Context) *DockerPullImages {
	var Images Images
	var DockerPullImages DockerPullImages
	if errA := ctx.ShouldBind(&Images); errA != nil {
		ctx.String(http.StatusUnprocessableEntity, `the body should be formA`)
	}
	requestimg := Images.Images

	DockerPullImages.Images = strings.Split(requestimg, ",")
	return &DockerPullImages
}
