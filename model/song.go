package model

type Song struct {
	Basic
	SingerId     int    `orm:"singer_id" json:"singerId"`
	Name         string `orm:"name" json:"name"`
	Introduction string `orm:"introduction" json:"introduction"`
	CreateTime   string `orm:"create_time" json:"create_time"` // 发行时间
	UpdateTime   string `orm:"update_time" json:"update_time"`
	Pic          string `orm:"pic" json:"pic"`
	Lyric        string `orm:"lyric" json:"lyric"`
	Url          string `orm:"url" json:"url"`
}

func (*Song) TableName() string {
	return "song"
}
