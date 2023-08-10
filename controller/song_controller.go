package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/constant"
	"go-music/response"
	"go-music/service"
	"net/url"
	"strconv"
)

type ISongController interface {
	SongList(c *gin.Context)        // 获取歌单列表
	SongListByStyle(c *gin.Context) // 根据风格搜索歌单
	ListSong(c *gin.Context)        // 获取歌单全部歌曲
	Song(c *gin.Context)            // 更加歌曲id获取歌曲
	SongListByTitle(c *gin.Context) // 根据标题搜索歌单
	SongByName(c *gin.Context)      // 根据歌名搜索歌曲
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
		response.Fail(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	// 解码URL参数
	decodedStyle, err := url.QueryUnescape(style)
	if err != nil {
		response.Fail(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	data := SC.songService.SongListByStyle(decodedStyle)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}

func (SC SongController) ListSong(c *gin.Context) {
	listId, ok := c.GetQuery("songListId")
	if !ok {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	id, err := strconv.Atoi(listId)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	data := SC.songService.ListSong(id)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}

func (SC SongController) Song(c *gin.Context) {
	songId, ok := c.GetQuery("id")
	if !ok {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	id, err := strconv.Atoi(songId)
	if err != nil {
		response.Error(c, nil, constant.PARAM_FAIL_PARSE)
		return
	}
	song := SC.songService.Song(id)
	response.Success(c, song, constant.REQUSEST_SUCCESS)
}

func (SC SongController) SongListByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	data := SC.songService.SongListByTitle(title)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}

func (SC SongController) SongByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		response.Error(c, nil, constant.PARAM_FAIL_GET)
		return
	}
	data := SC.songService.SongByName(name)
	response.Success(c, data, constant.REQUSEST_SUCCESS)
}
