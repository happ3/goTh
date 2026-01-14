package models

type User struct {
	Id             int    `gorm:"column:id" json:"id"`
	Username       string `gorm:"column:username" form:"username" json:"username"`
	Password       string `gorm:"column:pwz" form:"password" json:"password"`
	Email          string `gorm:"column:email" json:"email"`
	Authentication string `gorm:"column:auth" json:"authentication"`
}

func (*User) TableName() string {
	return "user"
}
