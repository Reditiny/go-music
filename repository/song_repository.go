package repository

import (
	"go-music/common"
	"go-music/model"
)

type ISongRepository interface {
	SongList() []*model.SongList
	SongListByStyle(style string) []*model.SongList
	ListSong(id int) []*model.ListSong
	Song(id int) []*model.Song
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

func (SR SongRepository) SongListByStyle(style string) []*model.SongList {
	var songList []*model.SongList
	// 模糊查询
	common.DB.Where("style LIKE ?", "%"+style+"%").Find(&songList)
	return songList
}

func (SR SongRepository) ListSong(id int) []*model.ListSong {
	var listSong []*model.ListSong
	common.DB.Where("song_list_id = ?", id).Find(&listSong)
	return listSong
}

func (SR SongRepository) Song(id int) []*model.Song {
	var song []*model.Song
	common.DB.Where("id = ?", id).Find(&song)
	return song
}
