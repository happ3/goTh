package controllers

import (
	"goTh/app/models"
	"goTh/app/models/response"
	"goTh/app/models/reuqest"
	"goTh/app/service"
	"time"

	"github.com/gin-gonic/gin"
)

/*
*
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。
*/
type CommentsController struct {
}

func (CommentsController) AddComments(c *gin.Context) {
	comment, err := reuqest.GetJsonToObj[models.Comments](c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
		return
	}
	if len(comment.Content) == 0 {
		response.FailMsg(c, "评论信息不能为空")
		return
	}
	if comment.PostId == 0 {
		response.FailMsg(c, "博客文章id不能为空")
		return
	}
	//取出session中当前用户
	userInterface, exists := c.Get("user")
	if !exists {
		return
	}
	user := userInterface.(*models.User)
	service.GteUserInfo(user)

	if user.Authentication != "1" {
		response.FailMsg(c, "只有已认证用户才能发表评论")
		return
	}
	comment.UserId = user.Id
	comment.CreateAt = time.Now()
	service.AddComments(&comment)
	response.SuccessMsg(c, "评论发表成功")
}

// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
func (CommentsController) FindCommentByPostId(c *gin.Context) {
	comments, err := reuqest.GetJsonToObj[models.Comments](c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
	}
	if comments.PostId == 0 {
		response.FailMsg(c, "文章的id不能为空")
		return
	}
	commentsArr := service.FindCommentByPostId(&comments)
	response.SuccessData(c, &commentsArr)
}
