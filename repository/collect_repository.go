package repository

import (
	"go-music/common"
	"go-music/model"
)

type ICollectRepository interface {
	Status(collect model.Collect) bool
	Add(collect model.Collect)
	Delete(collect model.Collect)
	GetCollection(userId int) []*model.Collect
}

type CollectRepository struct {
}

func NewCollectRepository() ICollectRepository {
	return CollectRepository{}
}

func (c CollectRepository) Status(collect model.Collect) bool {
	common.DB.Where("user_id = ?", collect.UserId).Where("song_id = ?", collect.SongId).Where("type = ?", collect.Type).Find(&collect)
	if collect.ID != 0 {
		return true
	} else {
		return false
	}
}

func (c CollectRepository) Add(collect model.Collect) {
	common.DB.Create(&collect)
}

func (c CollectRepository) Delete(collect model.Collect) {
	common.DB.Where("user_id = ?", collect.UserId).Where("song_id = ?", collect.SongId).First(&collect)
	common.DB.Where("id = ?", collect.ID).Delete(&collect)
}

func (c CollectRepository) GetCollection(userId int) []*model.Collect {
	var collects []*model.Collect
	common.DB.Where("user_id = ?", userId).Find(&collects)
	return collects
}
