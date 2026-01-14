package service

import (
	"goTh/app/models"
	"goTh/mysqlConfig"
	"goTh/util"
)

func AddComments(c *models.Comments) {
	keys := util.GetKeys(c)
	mysqlConfig.DB.Select(keys).Save(c)
}

func GetComments(c *models.Comments) {
	mysqlConfig.DB.Find(c)
}

func DeleteComments(c *models.Comments) {
	mysqlConfig.DB.Delete(c)
}
func UpdateComments(c *models.Comments) {
	mysqlConfig.DB.Model(c).Where("id=?", c.Id).Update("content", c.Content)
}

func FindCommentByPostId(c *models.Comments) []models.Comments {
	commentsArr := []models.Comments{}
	mysqlConfig.DB.Where("postId=?", c.PostId).Find(&commentsArr)
	return commentsArr
}
