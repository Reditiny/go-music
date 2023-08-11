package repository

import (
	"go-music/common"
	"go-music/model"
)

type ISingerRepository interface {
	Singers() []*model.Singer
	SingersBySex(sex int) []*model.Singer
	SingersById(singerId int) []*model.Song
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

func (SR SingerRepository) SingersBySex(sex int) []*model.Singer {
	var singers []*model.Singer
	common.DB.Where("sex = ?", sex).Find(&singers)
	return singers
}

func (SR SingerRepository) SingersById(singerId int) []*model.Song {
	var song []*model.Song
	common.DB.Where("singer_id = ?", singerId).Find(&song)
	return song
}
