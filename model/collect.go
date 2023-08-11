package model

type Collect struct {
	Basic
	UserId     int `orm:"user_id" json:"userId"`
	Type       int `orm:"type" json:"type"`
	SongId     int `orm:"song_id" json:"songId"`
	SongListId int `orm:"song_list_id" json:"songListId"`
}

func (*Collect) TableName() string {
	return "collect"
}
