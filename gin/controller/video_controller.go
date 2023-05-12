package controller

import (
	"barbz.dev/gin/entity"
	"barbz.dev/gin/service"
	"github.com/gin-gonic/gin"
	"log"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)

	if err != nil {
		log.Fatal(err)
		return entity.Video{}
	}
	c.service.Save(video)

	return video
}
