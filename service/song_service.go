package service

import (
	"encoding/json"
	"go-music/common"
	"go-music/constant"
	"go-music/model"
	"go-music/repository"
)

type ISongService interface {
	SongList() []*model.SongList
	SongListByStyle(style string) []*model.SongList
	ListSong(id int) []*model.ListSong
	Song(id int) []*model.Song
	SongListByTitle(title string) []*model.SongList
	SongByName(name string) []*model.Song
}

type SongService struct {
	songRepository repository.ISongRepository
}

func NewSongService() ISongService {
	return SongService{songRepository: repository.NewSongRepository()}
}

func (SS SongService) SongList() []*model.SongList {
	songList, err := common.RedisClient.Get(common.Ctx, constant.INDEX_SONG_LIST).Bytes()
	var data []*model.SongList
	if err != nil {
		data = SS.songRepository.SongList()
		// 序列化为 JSON 字节流
		jsonData, _ := json.Marshal(data)
		common.RedisClient.Set(common.Ctx, constant.INDEX_SONG_LIST, jsonData, -1)
	} else {
		json.Unmarshal(songList, &data)
	}
	return data
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

func (SS SongService) SongListByTitle(title string) []*model.SongList {
	return SS.songRepository.SongListByTitle(title)
}

func (SS SongService) SongByName(name string) []*model.Song {
	return SS.songRepository.SongByName(name)
}
