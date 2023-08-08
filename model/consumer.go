package model

import "gorm.io/gorm"

type Consumer struct {
	gorm.Model
	Username     string `orm:"username" json:"username"`
	Password     string `orm:"password" json:"password"`
	Sex          int    `orm:"sex" json:"sex"`
	PhoneNum     string `orm:"phone_num" json:"phone_num"`
	Email        string `orm:"email" json:"email"`
	Birth        string `orm:"birth" json:"birth"`
	Introduction string `orm:"introduction" json:"introduction"`
	Location     string `orm:"location" json:"location"`
	Avator       string `orm:"avator" json:"avator"`
	CreateTime   string `orm:"create_time" json:"create_time"`
	UpdateTime   string `orm:"update_time" json:"update_time"`
}

func (*Consumer) TableName() string {
	return "consumer"
}
