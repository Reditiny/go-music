package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
	_ "go-music/controller"
)

func InitCollect(r *gin.RouterGroup) {
	router := r.Group("/collection")
	collectController := controller.NewCollectController()
	{
		router.POST("/status", collectController.Status)
		router.POST("/add", collectController.Add)
		router.DELETE("/delete", collectController.Delete)
		router.GET("/detail", collectController.GetCollection)
	}
}
