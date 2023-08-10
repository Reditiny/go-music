package model

type Banner struct {
	Basic
	Pic string `orm:"pic" json:"pic"`
}

func (*Banner) TableName() string {
	return "banner"
}
