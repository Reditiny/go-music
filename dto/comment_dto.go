package dto

type CommentDto struct {
	UserId     int    `orm:"user_id" json:"userId"`
	SongId     int    `orm:"song_id" json:"songId"`
	SongListId string `orm:"song_list_id" json:"songListId"`
	Content    string `orm:"content" json:"content"`
	Type       int    `orm:"type" json:"nowType"`
}
