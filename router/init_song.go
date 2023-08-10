package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func InitSong(r *gin.RouterGroup) {
	router := r.Group("/")
	songController := controller.NewSongController()
	{
		router.GET("/songList", songController.SongList)
		router.GET("/songList/style/detail", songController.SongListByStyle)
		router.GET("/listSong/detail", songController.ListSong)
		router.GET("/song/detail", songController.Song)
	}
}
