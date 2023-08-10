package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type ISingerController interface {
	Singers(c *gin.Context)
	SingersBySex(c *gin.Context)
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

func (SC SingerController) SingersBySex(c *gin.Context) {
	sex, ok := c.GetQuery("sex")
	sexInt, err := strconv.Atoi(sex)
	if !ok || err != nil {
		response.Fail(c, nil, "参数错误")
	} else {
		data := SC.singerService.SingersBySex(sexInt)
		response.Success(c, data, "歌手列表")
	}

}
