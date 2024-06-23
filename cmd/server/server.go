package main

import (
	"github.com/gin-gonic/gin"
	"operator-dev/docker-image-download/internal/app/server"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.POST("/api/img/pull", server.RunPull)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	//r.POST("/api/img/test", server.JsonData)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
