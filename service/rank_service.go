package service

import (
	"go-music/dto"
	"go-music/model"
	"go-music/repository"
)

type IRankService interface {
	AvgScore(song int) int
	MyScore(consumer int, song int) *model.RankList
	AddScore(score dto.RankDto)
}

type RankService struct {
	rankRepository repository.IRankRepository
}

func NewRankService() IRankService {
	return RankService{rankRepository: repository.NewRankRepository()}
}

func (RS RankService) AvgScore(song int) int {
	list := RS.rankRepository.AvgScore(song)
	score := 0
	sum := 0
	count := 0
	for _, rank := range list {
		sum += rank.Score
		count++
	}
	if count != 0 {
		score = sum / count
	}
	return score
}

func (RS RankService) MyScore(consumer int, song int) *model.RankList {
	return RS.rankRepository.MyScore(consumer, song)
}

func (RS RankService) AddScore(score dto.RankDto) {
	rankList := model.RankList{
		ConsumerId: score.ConsumerId,
		SongListId: score.SongListId,
		Score:      score.Score,
	}
	RS.rankRepository.AddScore(rankList)
}
