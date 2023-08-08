package repository

import (
	"go-music/common"
	"go-music/model"
)

type ISongRepository interface {
	SongList() []*model.SongList
}

type SongRepository struct {
}

func NewSongRepository() ISongRepository {
	return SongRepository{}
}

func (SR SongRepository) SongList() []*model.SongList {
	var songList []*model.SongList
	common.DB.Find(&songList)
	return songList
}
