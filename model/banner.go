package model

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Pic string `orm:"pic" json:"pic"`
}

func (*Banner) TableName() string {
	return "banner"
}
