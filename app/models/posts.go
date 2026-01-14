package models

import "time"

type Post struct {
	Id       int        `gorm:"column:id" json:"id"`
	Title    string     `gorm:"column:title" json:"title"`
	Content  string     `gorm:"column:content"  json:"content"`
	UserId   int        `gorm:"column:userId" json:"user_id"`
	CreateAt time.Time  `gorm:"column:createAt" json:"createAt"`
	UpdateAt *time.Time `gorm:"column:updateAt" json:"updateAt,omitempty"` // 使用指针类型
}

/*
UpdateAt *time.Time 改为指针类型
当值为 nil 时，JSON 中不显示该字段
仅当有实际值时才显示 update_at
现在 update_at 字段只有在有值时才会出现在 JSON 输出中！
*
*/
func (*Post) TableName() string {
	return "posts"
}
