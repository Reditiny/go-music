package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId     int    `orm:"user_id" json:"user_id"`
	SongId     int    `orm:"song_id" json:"song_id"`
	SongListId int    `orm:"song_list_id" json:"song_list_id"`
	Content    string `orm:"content" json:"content"`
	CreateTime string `orm:"create_time" json:"create_time"`
	Type       int    `orm:"type" json:"type"`
	Up         int    `orm:"up" json:"up"`
}

func (*Comment) TableName() string {
	return "comment"
}
