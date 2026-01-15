package service

import (
	"goTh/app/common"
	"goTh/app/dto"
	"goTh/app/models"
	"goTh/mysqlConfig"
	"goTh/util"

	"gorm.io/gorm"
)

func PagePost(postsDto dto.PostsDto, offset int, pageSize int) common.PageResp {
	var posts []models.Post
	var total int64

	mysqlConfig.DB.Scopes(buildQuery(postsDto)).Model(&models.Post{}).Count(&total)
	mysqlConfig.DB.Scopes(buildQuery(postsDto)).Model(&models.Post{}).Offset(offset).Limit(pageSize).Find(&posts)
	return common.PageResp{}.PageResult(offset, pageSize, total, posts)

}

func NewPagePost(postsDto dto.PostsDto, paging *dto.Paging) common.PageResp {
	var posts []models.Post
	var total int64

	mysqlConfig.DB.Scopes(buildQuery(postsDto)).Model(&models.Post{}).Count(&total)
	mysqlConfig.DB.Scopes(buildQuery(postsDto)).Model(&models.Post{}).Offset(paging.Offset).Limit(paging.PageSize).Find(&posts)
	return common.PageResp{}.PageResult(paging.Offset, paging.PageSize, total, posts)
}

func buildQuery(postsDto dto.PostsDto) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(postsDto.Content) != 0 {
			db = db.Where("content = ?", postsDto.Content)
		}
		if len(postsDto.Title) != 0 {
			db = db.Where("title LIKE ?", "%"+postsDto.Title+"%")
		}
		return db
	}
}

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
