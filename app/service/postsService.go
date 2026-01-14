package service

import (
	"goTh/app/models"
	"goTh/mysqlConfig"
	"goTh/util"
)

/*
*
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
*/
func AddPosts(p *models.Post) {
	keys := util.GetKeys(p)
	mysqlConfig.DB.Select(keys).Save(p)
}

/*
*
实现文章的更新功能，只有文章的作者才能更新自己的文章。
*/
func UpdatePosts(p *models.Post) {
	post := models.Post{}
	mysqlConfig.DB.Model(&post).Where("id", p.Id).Update("title", p.Title).Update("content", p.Content)
}

func Del(p *models.Post) {
	mysqlConfig.DB.Where("id", p.Id).Delete(p)
}

func GetPost(p *models.Post) {
	mysqlConfig.DB.Where("id = ?", p.Id).Find(p)
}
