package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/response"
	"go-music/service"
)

type IBannerController interface {
	GetAllBanner(c *gin.Context) // 获取所有轮播图接口
}

type BannerController struct {
	bannerService service.IBannerService
}

func NewBannerController() IBannerController {
	return BannerController{bannerService: service.NewBannerService()}
}

func (BC BannerController) GetAllBanner(c *gin.Context) {
	data := BC.bannerService.GetAllBanner()
	response.Success(c, data, "成功获取轮播图与")
}
