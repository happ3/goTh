package models

import "time"

type Comments struct {
	Id       int       `json:"id" gorm:"column:id"`
	Content  string    `json:"content" gorm:"column:content"`
	UserId   int       `json:"userId" gorm:"column:userId"`
	PostId   int       `json:"PostId" gorm:"column:postId"`
	CreateAt time.Time `json:"createAt" gorm:"column:createAt"`
	UpdateAt time.Time `json:"updateAt" gorm:"column:updateAt"`
}

func (*Comments) TableName() string {
	return "comments"
}
