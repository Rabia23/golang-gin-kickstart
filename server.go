package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Rabia23/golang-gin-kickstart/controller"
	"github.com/Rabia23/golang-gin-kickstart/middlewares"
	"github.com/Rabia23/golang-gin-kickstart/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Valid Input!"})
		}
	})

	server.Run(":8080")
}
