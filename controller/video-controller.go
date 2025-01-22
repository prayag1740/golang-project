package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prayag1740/golang-project/entity"
	"github.com/prayag1740/golang-project/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type VideoControllerImpl struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &VideoControllerImpl{service: service}
}

func (controller *VideoControllerImpl) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *VideoControllerImpl) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return controller.service.Save(video)
}
