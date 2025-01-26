package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/prayag1740/golang-project/entity"
	"github.com/prayag1740/golang-project/service"
	"github.com/prayag1740/golang-project/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
}

type VideoControllerImpl struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {

	validate = validator.New()
	validate.RegisterValidation("title_contains_cool", validators.ValidateCoolTitle)

	return &VideoControllerImpl{service: service}
}

func (controller *VideoControllerImpl) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *VideoControllerImpl) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}
	err = validate.Struct(video)
	if err != nil {
		return entity.Video{}, err
	}
	controller.service.Save(video)
	return video, nil
}
