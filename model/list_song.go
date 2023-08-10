package model

type ListSong struct {
	Basic
	SongId     int `orm:"song_id" json:"songId"`
	SongListId int `orm:"song_list_id" json:"song_list_id"`
}

func (*ListSong) TableName() string {
	return "list_song"
}
