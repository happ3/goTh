package service

import (
	"fmt"
	"goTh/app/models"
	"goTh/mysqlConfig"

	"golang.org/x/crypto/bcrypt"
)

func GteUserInfo(u *models.User) {
	if u.Id != 0 {
		mysqlConfig.DB.Where("id = ?", u.Id).First(u)
	} else if len(u.Username) != 0 {
		mysqlConfig.DB.Where("username = ?", u.Username).First(u)
	}

}

func Add(u *models.User) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	mysqlConfig.DB.Create(u)
}

func CheckUserInfo(u *models.User) bool {
	copyUser := u
	mysqlConfig.DB.Where("username = ?", copyUser.Username).Find(&copyUser)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if u.Password == string(hashedPassword) {
		return true
	}
	return false
}

/*
*
更新用户为已认证
*/
func UpdateUserInfo(u *models.User) error {
	result := mysqlConfig.DB.Model(&models.User{}).Where("id = ?", u.Id).Update("auth", u.Authentication)
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户id不存在: %d", u.Id)
	}
	/**
	方法2
	mysqlConfig.DB.Table("user").Where("id = ?", u.Id).Update("auth", u.Authentication)

	方法3
	updates := map[string]interface{}{
		"authentication": "1",
	}
	mysqlConfig.DB.Model(&models.User{}).Where("id", u.Id).Updates(updates)

	方法4
	newUser := &models.User{
		Id:             u.Id,
		Authentication: "1",
	}
	mysqlConfig.DB.Select("authentication").Save(&newUser)
	**/
	return nil
}
