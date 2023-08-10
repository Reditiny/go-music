package service

import (
	"go-music/model"
	"go-music/repository"
)

type ISongService interface {
	SongList() []*model.SongList
	SongListByStyle(style string) []*model.SongList
	ListSong(id int) []*model.ListSong
	Song(id int) []*model.Song
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

func (SS SongService) SongListByStyle(style string) []*model.SongList {
	return SS.songRepository.SongListByStyle(style)
}

func (SS SongService) ListSong(id int) []*model.ListSong {
	return SS.songRepository.ListSong(id)
}

func (SS SongService) Song(id int) []*model.Song {
	return SS.songRepository.Song(id)
}
