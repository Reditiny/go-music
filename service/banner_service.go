package service

import (
	"go-music/model"
	"go-music/repository"
	"go-music/utils"
	"go-music/vo"
)

type IBannerService interface {
	GetAllBanner() []*vo.BannerVo
}

type BannerService struct {
	bannerRepository repository.IBannerRepository
}

func NewBannerService() IBannerService {
	return BannerService{bannerRepository: repository.NewBannerRepository()}
}

func (BS BannerService) GetAllBanner() []*vo.BannerVo {
	banners := BS.bannerRepository.GetAllBanner()
	data := utils.CopyList[*model.Banner, *vo.BannerVo](banners)
	return data
}
