package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/response"
	"go-music/service"
	"net/url"
	"strconv"
)

type ISongController interface {
	SongList(c *gin.Context)
	SongListByStyle(c *gin.Context)
	ListSong(c *gin.Context)
	Song(c *gin.Context)
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

func (SC SongController) SongListByStyle(c *gin.Context) {
	// 获取查询参数
	style, ok := c.GetQuery("style")
	if !ok {
		response.Fail(c, nil, "参数获取失败")
		return
	}
	// 解码URL参数
	decodedStyle, err := url.QueryUnescape(style)
	if err != nil {
		response.Fail(c, nil, "参数解析失败")
		return
	}
	data := SC.songService.SongListByStyle(decodedStyle)
	response.Success(c, data, "请求成功")
}

func (SC SongController) ListSong(c *gin.Context) {
	listId, ok := c.GetQuery("songListId")
	if !ok {
		response.Error(c, nil, "参数获取失败")
		return
	}
	id, err := strconv.Atoi(listId)
	if err != nil {
		response.Error(c, nil, "参数不合法")
		return
	}
	data := SC.songService.ListSong(id)
	response.Success(c, data, "成功")
}

func (SC SongController) Song(c *gin.Context) {
	songId, ok := c.GetQuery("id")
	if !ok {
		response.Error(c, nil, "参数获取失败")
		return
	}
	id, err := strconv.Atoi(songId)
	if err != nil {
		response.Error(c, nil, "参数不合法")
		return
	}
	song := SC.songService.Song(id)
	response.Success(c, song, "成功")
}
