package main

import (
	"github.com/Rabia23/golang-gin-kickstart/controller"
	"github.com/Rabia23/golang-gin-kickstart/service"
	"github.com/gin-gonic/gin"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
