package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func InitRank(r *gin.RouterGroup) {
	router := r.Group("/")
	rankController := controller.NewRankController()
	{
		router.GET("/rankList", rankController.AvgScore)
		router.GET("/rankList/user", rankController.MyScore)
		router.POST("/rankList/add", rankController.AddScore)
	}
}
