package model

type UserSupport struct {
	Basic
	CommentId int    `orm:"comment_id" json:"comment_id"`
	UserId    string `orm:"user_id" json:"user_id"`
}

func (*UserSupport) TableName() string {
	return "user_support"
}
