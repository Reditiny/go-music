package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/constant"
	"go-music/dto"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type IRankController interface {
	AvgScore(c *gin.Context)
	MyScore(c *gin.Context)
	AddScore(c *gin.Context)
}

type RankController struct {
	rankService service.IRankService
}

func NewRankController() IRankController {
	return RankController{rankService: service.NewRankService()}
}

func (RC RankController) MyScore(c *gin.Context) {
	var songListId, consumerId string
	songListId = c.Query("songListId")
	consumerId = c.Query("consumerId")
	if songListId == "" || consumerId == "" {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	song, err1 := strconv.Atoi(songListId)
	consumer, err2 := strconv.Atoi(consumerId)
	if err1 != nil || err2 != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	rank := RC.rankService.MyScore(consumer, song)
	if rank.ID == 0 {
		response.Error(c, nil, "未评价")
	} else {
		response.Success(c, rank.Score, constant.REQUSEST_SUCCESS)
	}
}

func (RC RankController) AvgScore(c *gin.Context) {
	var songListId string
	songListId = c.Query("songListId")
	if songListId == "" {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	song, err := strconv.Atoi(songListId)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	score := RC.rankService.AvgScore(song)
	response.Success(c, score, constant.REQUSEST_SUCCESS)
}

func (RC RankController) AddScore(c *gin.Context) {
	var rankScore dto.RankDto
	err := c.ShouldBind(&rankScore)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
	}
	RC.rankService.AddScore(rankScore)
	response.Success(c, nil, constant.REQUSEST_SUCCESS)
}
