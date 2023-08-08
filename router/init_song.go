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
	}
}
