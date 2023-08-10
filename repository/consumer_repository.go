package repository

import (
	"go-music/common"
	"go-music/dto"
	"go-music/model"
)

type IConsumerRepository interface {
	Login(user dto.LoginDto) []*model.Consumer
	Add(consumer model.Consumer) bool
	GetById(id int) []*model.Consumer
}

type ConsumerRepository struct {
}

func NewConsumerRepository() IConsumerRepository {
	return ConsumerRepository{}
}

func (c ConsumerRepository) Login(user dto.LoginDto) []*model.Consumer {
	var consumer []*model.Consumer
	common.DB.Where("username = ?", user.Username).Where("password = ?", user.Password).Find(&consumer)
	return consumer
}

func (c ConsumerRepository) Add(consumer model.Consumer) bool {
	var one model.Consumer
	common.DB.Where("username = ?", consumer.Username).Find(&one)
	if one.Username == consumer.Username {
		return false
	} else {
		common.DB.Create(&consumer)
		return true
	}
}

func (c ConsumerRepository) GetById(id int) []*model.Consumer {
	var consumer []*model.Consumer
	common.DB.Where("id = ?", id).Find(&consumer)
	return consumer
}
