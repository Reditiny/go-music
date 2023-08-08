package model

import "gorm.io/gorm"

type Collect struct {
	gorm.Model
	UserId     int    `orm:"user_id" json:"user_id"`
	Type       int    `orm:"type" json:"type"`
	SongId     int    `orm:"song_id" json:"song_id"`
	SongListId int    `orm:"song_list_id" json:"song_list_id"`
	CreateTime string `orm:"create_time" json:"create_time"`
}

func (*Collect) TableName() string {
	return "collect"
}
