package model

type SongList struct {
	Basic
	Title        string `orm:"title" json:"title"`
	Pic          string `orm:"pic" json:"pic"`
	Introduction string `orm:"introduction" json:"introduction"`
	Style        string `orm:"style" json:"style"`
}

func (*SongList) TableName() string {
	return "song_list"
}
