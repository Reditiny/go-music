package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/constant"
	"go-music/dto"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type ICommentController interface {
	SongListComment(c *gin.Context)
	Add(c *gin.Context)
}

type CommentController struct {
	commentService service.ICommentService
}

func NewCommentController() ICommentController {
	return CommentController{commentService: service.NewCommentService()}
}

func (CC CommentController) SongListComment(c *gin.Context) {
	songListId := c.Query("songListId")
	if songListId == "" {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	songList, err := strconv.Atoi(songListId)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	data := CC.commentService.SongListComment(songList)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}

func (CC CommentController) Add(c *gin.Context) {
	var comment dto.CommentDto
	err := c.ShouldBind(&comment)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
	}
	CC.commentService.AddComment(comment)
	response.Success(c, nil, constant.REQUSEST_SUCCESS)
}
