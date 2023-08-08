package model

import "gorm.io/gorm"

type UserSupport struct {
	gorm.Model
	CommentId int    `orm:"comment_id" json:"comment_id"`
	UserId    string `orm:"user_id" json:"user_id"`
}

func (*UserSupport) TableName() string {
	return "user_support"
}
