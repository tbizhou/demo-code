package main

import (
	"github.com/demo-code/internal/app/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/img/pull", server.RunPull)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
