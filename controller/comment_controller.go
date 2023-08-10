package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/response"
	"go-music/service"
	"strconv"
)

type ICommentController interface {
	SongListComment(c *gin.Context)
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
		response.Error(c, nil, "参数缺失")
		return
	}
	songList, err := strconv.Atoi(songListId)
	if err != nil {
		response.Error(c, nil, "无效参数")
		return
	}
	data := CC.commentService.SongListComment(songList)
	response.Success(c, data, "获取成功")
}
