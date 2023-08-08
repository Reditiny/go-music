package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/response"
	"go-music/service"
)

type ISingerController interface {
	Singers(c *gin.Context)
}

type SingerController struct {
	singerService service.ISingerService
}

func NewSingerController() ISingerController {
	return SingerController{singerService: service.NewSingerService()}
}

func (SC SingerController) Singers(c *gin.Context) {
	singers := SC.singerService.Singers()
	response.Success(c, singers, "歌手列表")
}
