package model

import "gorm.io/gorm"

type ListSong struct {
	gorm.Model
	SongId     int `orm:"song_id" json:"song_id"`
	SongListId int `orm:"song_list_id" json:"song_list_id"`
}

func (*ListSong) TableName() string {
	return "list_song"
}
