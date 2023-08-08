package repository

import (
	"go-music/common"
	"go-music/model"
)

type IBannerRepository interface {
	GetAllBanner() []*model.Banner
}

type BannerRepository struct {
}

func NewBannerRepository() IBannerRepository {
	return BannerRepository{}
}

func (BR BannerRepository) GetAllBanner() []*model.Banner {
	var banners []*model.Banner
	common.DB.Find(&banners)
	return banners
}
