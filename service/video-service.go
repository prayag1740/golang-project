package service

import "github.com/prayag1740/golang-project/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type VideoServiceImpl struct {
	videos []entity.Video
}

func New() VideoService {
	return &VideoServiceImpl{}
}

func (service *VideoServiceImpl) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *VideoServiceImpl) FindAll() []entity.Video {
	return service.videos
}
