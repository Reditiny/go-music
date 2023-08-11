package service

import (
	"go-music/model"
	"go-music/repository"
)

type ISingerService interface {
	Singers() []*model.Singer
	SingersBySex(sex int) []*model.Singer
	SingersById(singerId int) []*model.Song
}

type SingerService struct {
	singerRepository repository.ISingerRepository
}

func NewSingerService() ISingerService {
	return SingerService{singerRepository: repository.NewSingerRepository()}
}

func (SS SingerService) Singers() []*model.Singer {
	singers := SS.singerRepository.Singers()
	return singers
}

func (SS SingerService) SingersBySex(sex int) []*model.Singer {
	return SS.singerRepository.SingersBySex(sex)
}

func (SS SingerService) SingersById(singerId int) []*model.Song {
	return SS.singerRepository.SingersById(singerId)
}
