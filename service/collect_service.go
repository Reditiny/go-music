package service

import (
	"go-music/dto"
	"go-music/model"
	"go-music/repository"
	"strconv"
)

type ICollectService interface {
	Status(status dto.CollectStatusDto) bool
	Add(status dto.CollectStatusDto)
	Delete(userid int, songId int)
	GetCollection(userId int) []*model.Collect
}

type CollectService struct {
	collectRepository repository.ICollectRepository
}

func NewCollectService() ICollectService {
	return CollectService{collectRepository: repository.NewCollectRepository()}
}

func (CS CollectService) Status(status dto.CollectStatusDto) bool {
	t, _ := strconv.Atoi(status.Type)
	collect := model.Collect{
		UserId: status.UserId,
		SongId: status.SongId,
		Type:   t,
	}
	return CS.collectRepository.Status(collect)
}

func (CS CollectService) Add(status dto.CollectStatusDto) {
	t, _ := strconv.Atoi(status.Type)
	collect := model.Collect{
		UserId: status.UserId,
		SongId: status.SongId,
		Type:   t,
	}
	CS.collectRepository.Add(collect)
}

func (CS CollectService) Delete(userId int, songId int) {
	collect := model.Collect{UserId: userId, SongId: songId}
	CS.collectRepository.Delete(collect)
}

func (CS CollectService) GetCollection(userId int) []*model.Collect {
	return CS.collectRepository.GetCollection(userId)
}
