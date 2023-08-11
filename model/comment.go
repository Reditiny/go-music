package model

type Comment struct {
	Basic
	UserId     int    `orm:"user_id" json:"userId"`
	SongId     int    `orm:"song_id" json:"song_id"`
	SongListId int    `orm:"song_list_id" json:"song_list_id"`
	Content    string `orm:"content" json:"content"`
	Type       int    `orm:"type" json:"type"`
	Up         int    `orm:"up" json:"up"`
}

func (*Comment) TableName() string {
	return "comment"
}
