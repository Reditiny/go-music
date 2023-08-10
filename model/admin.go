package model

type Admin struct {
	Basic
	Name     string `orm:"name" json:"name"`
	Password string `orm:"password" json:"password"`
}

func (*Admin) TableName() string {
	return "admin"
}
