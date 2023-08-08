package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func InitSinger(r *gin.RouterGroup) {
	router := r.Group("/")
	singerController := controller.NewSingerController()
	{
		router.GET("/singer", singerController.Singers)
	}
}
