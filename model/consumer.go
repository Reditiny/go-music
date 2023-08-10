package model

import (
	"time"
)

type Consumer struct {
	Basic
	Username     string    `orm:"username" json:"username"`
	Password     string    `orm:"password" json:"password"`
	Sex          int       `orm:"sex" json:"sex"`
	PhoneNum     string    `orm:"phone_num" json:"phone_num"`
	Email        string    `orm:"email" json:"email"`
	Birth        time.Time `orm:"birth" json:"birth"`
	Introduction string    `orm:"introduction" json:"introduction"`
	Location     string    `orm:"location" json:"location"`
	Avator       string    `orm:"avator" json:"avator"`
}

func (*Consumer) TableName() string {
	return "consumer"
}
