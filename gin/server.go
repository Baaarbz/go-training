package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})

	server.Run(":8080")
}
