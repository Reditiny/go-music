package model

import "gorm.io/gorm"

type RankList struct {
	gorm.Model
	SongListId int `orm:"song_list_id" json:"song_list_id"`
	ConsumerId int `orm:"consumer_id" json:"consumer_id"`
	Score      int `orm:"score" json:"score"`
}

func (*RankList) TableName() string {
	return "rank_list"
}
