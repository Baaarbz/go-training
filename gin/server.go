package main

import (
	"barbz.dev/gin/controller"
	"barbz.dev/gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "UP",
		})
	})

	server.GET("/v1/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/v1/videos", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(":8080")
}
