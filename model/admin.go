package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `orm:"name" json:"name"`
	Password string `orm:"password" json:"password"`
}

func (*Admin) TableName() string {
	return "admin"
}
