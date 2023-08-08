package service

import (
	"go-music/model"
	"go-music/repository"
)

type ISingerService interface {
	Singers() []*model.Singer
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
