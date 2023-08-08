package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/response"
	"go-music/service"
)

type ISongController interface {
	SongList(c *gin.Context)
}

type SongController struct {
	songService service.ISongService
}

func NewSongController() ISongController {
	return SongController{songService: service.NewSongService()}
}

func (SC SongController) SongList(c *gin.Context) {
	songList := SC.songService.SongList()
	response.Success(c, songList, "歌曲列表")
}
