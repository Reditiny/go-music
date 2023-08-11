package repository

import (
	"go-music/common"
	"go-music/model"
)

type IRankRepository interface {
	AvgScore(song int) []*model.RankList
	MyScore(consumer int, song int) *model.RankList
	AddScore(score model.RankList)
}

type RankRepository struct {
}

func NewRankRepository() IRankRepository {
	return RankRepository{}
}

func (r RankRepository) AvgScore(song int) []*model.RankList {
	var list []*model.RankList
	common.DB.Where("song_list_id = ?", song).Find(&list)
	return list
}

func (r RankRepository) MyScore(consumer int, song int) *model.RankList {
	var list *model.RankList
	common.DB.Where("consumer_id = ?", consumer).Where("song_list_id = ?", song).First(&list)
	return list
}

func (r RankRepository) AddScore(score model.RankList) {
	common.DB.Create(&score)
}
