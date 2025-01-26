package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prayag1740/golang-project/controller"
	"github.com/prayag1740/golang-project/middleware"
	"github.com/prayag1740/golang-project/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	setupLogOutput()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/get-video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/post-video", func(ctx *gin.Context) {
		newVideo, err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, newVideo)
	})

	server.Run(":8080")
}
