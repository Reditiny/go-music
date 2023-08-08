package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func InitBanner(r *gin.RouterGroup) {
	router := r.Group("/banner")
	bannerController := controller.NewBannerController()
	{
		router.GET("/getAllBanner", bannerController.GetAllBanner)
	}
}
