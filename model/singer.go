package model

type Singer struct {
	Basic
	Name         string `orm:"name" json:"name"`
	Sex          int    `orm:"sex" json:"sex"`
	Pic          string `orm:"pic" json:"pic"`
	Birth        string `orm:"birth" json:"birth"`
	Location     string `orm:"location" json:"location"`
	Introduction string `orm:"introduction" json:"introduction"`
}

func (*Singer) TableName() string {
	return "singer"
}
