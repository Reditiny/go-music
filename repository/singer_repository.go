package repository

import (
	"go-music/common"
	"go-music/model"
)

type ISingerRepository interface {
	Singers() []*model.Singer
}

type SingerRepository struct {
}

func NewSingerRepository() ISingerRepository {
	return SingerRepository{}
}

func (SR SingerRepository) Singers() []*model.Singer {
	var singers []*model.Singer
	common.DB.Find(&singers)
	return singers
}
