package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func InitUser(r *gin.RouterGroup) {
	router := r.Group("/user")
	consumerController := controller.NewConsumerController()
	{
		router.POST("/login/status", consumerController.Login)
		router.POST("/add", consumerController.Add)
		router.GET("/detail", consumerController.GetById)
		router.POST("/update", consumerController.Update)
	}
}
