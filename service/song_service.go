package service

import (
	"go-music/model"
	"go-music/repository"
)

type ISongService interface {
	SongList() []*model.SongList
}

type SongService struct {
	songRepository repository.ISongRepository
}

func NewSongService() ISongService {
	return SongService{songRepository: repository.NewSongRepository()}
}

func (SS SongService) SongList() []*model.SongList {
	songList := SS.songRepository.SongList()
	return songList
}
