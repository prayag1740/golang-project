package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prayag1740/golang-project/controller"
	"github.com/prayag1740/golang-project/service"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func main() {
	server := gin.Default()

	server.GET("/get-video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/post-video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
