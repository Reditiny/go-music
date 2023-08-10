package service

import (
	"go-music/dto"
	"go-music/model"
	"go-music/repository"
)

type IConsumerService interface {
	Login(user dto.LoginDto) []*model.Consumer
	Add(consumer model.Consumer) bool
	GetById(id int) []*model.Consumer
}

type ConsumerService struct {
	consumerRepository repository.IConsumerRepository
}

func NewConsumerService() IConsumerService {
	return ConsumerService{consumerRepository: repository.NewConsumerRepository()}
}

func (CS ConsumerService) Login(user dto.LoginDto) []*model.Consumer {
	consumer := CS.consumerRepository.Login(user)
	return consumer
}

func (CS ConsumerService) Add(consumer model.Consumer) bool {
	return CS.consumerRepository.Add(consumer)
}

func (CS ConsumerService) GetById(id int) []*model.Consumer {
	return CS.consumerRepository.GetById(id)
}
