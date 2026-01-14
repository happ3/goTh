package models

import "time"

type Post struct {
	Id       int       `gorm:"column:id" json:"id"`
	Title    string    `gorm:"column:title" json:"title"`
	Content  string    `gorm:"column:content"  json:"content"`
	UserId   int       `gorm:"column:userId" json:"user_id"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
	UpdateAt time.Time `gorm:"column:updateAt" json:"updateAt"`
}

func (*Post) TableName() string {
	return "posts"
}
