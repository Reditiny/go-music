package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/constant"
	"go-music/dto"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type ICollectController interface {
	Status(c *gin.Context)
	Add(c *gin.Context)
	Delete(c *gin.Context)
	GetCollection(c *gin.Context)
}

type CollectController struct {
	collectService service.ICollectService
}

func NewCollectController() ICollectController {
	return CollectController{collectService: service.NewCollectService()}
}

func (CC CollectController) Status(c *gin.Context) {
	var collectStatus dto.CollectStatusDto
	err := c.ShouldBind(&collectStatus)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
	}
	ok := CC.collectService.Status(collectStatus)
	if ok {
		response.Success(c, ok, constant.STATUS_COLLECTED)
	} else {
		response.Success(c, ok, constant.STATUS_UNCOLLECTED)
	}
}

func (CC CollectController) Add(c *gin.Context) {
	var collectStatus dto.CollectStatusDto
	err := c.ShouldBind(&collectStatus)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
	}
	CC.collectService.Add(collectStatus)
	response.Success(c, nil, constant.REQUSEST_SUCCESS)
}

func (CC CollectController) Delete(c *gin.Context) {
	userIdStr := c.Query("userId")
	songIdStr := c.Query("songId")
	userId, err1 := strconv.Atoi(userIdStr)
	songId, err2 := strconv.Atoi(songIdStr)
	if err1 != nil || err2 != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	CC.collectService.Delete(userId, songId)
	response.Success(c, nil, constant.REQUSEST_SUCCESS)
}

func (CC CollectController) GetCollection(c *gin.Context) {
	userIdStr := c.Query("userId")
	userId, err1 := strconv.Atoi(userIdStr)
	if err1 != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	data := CC.collectService.GetCollection(userId)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}
