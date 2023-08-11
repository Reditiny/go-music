package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/constant"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type ISingerController interface {
	Singers(c *gin.Context)
	SingersBySex(c *gin.Context)
	SingersById(c *gin.Context)
}

type SingerController struct {
	singerService service.ISingerService
}

func NewSingerController() ISingerController {
	return SingerController{singerService: service.NewSingerService()}
}

func (SC SingerController) Singers(c *gin.Context) {
	singers := SC.singerService.Singers()
	response.Success(c, singers, constant.REQUSEST_SUCCESS)
}

func (SC SingerController) SingersBySex(c *gin.Context) {
	sex, ok := c.GetQuery("sex")
	sexInt, err := strconv.Atoi(sex)
	if !ok || err != nil {
		response.Fail(c, nil, constant.PARAM_FAIL_GET)
	} else {
		data := SC.singerService.SingersBySex(sexInt)
		response.Success(c, data, constant.REQUSEST_SUCCESS)
	}
}

func (SC SingerController) SingersById(c *gin.Context) {
	singerIdStr := c.Query("singerId")
	singerId, err := strconv.Atoi(singerIdStr)
	if err != nil {
		response.Fail(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	data := SC.singerService.SingersById(singerId)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}
