package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func InitComment(r *gin.RouterGroup) {
	router := r.Group("/")
	commentController := controller.NewCommentController()
	{
		router.GET("/comment/songList/detail", commentController.SongListComment)
	}
}
