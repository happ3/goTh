package controllers

import (
	"goTh/app/dto"
	"goTh/app/models"
	"goTh/app/models/response"
	"goTh/app/models/reuqest"
	"goTh/app/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @title 博客API文档
// @version 1.0
// @description 博客系统的API接口文档
// @host localhost:8080
// @BasePath /api/v1
type PostsController struct {
	BaseController
}

func (PostsController) PagePost(c *gin.Context) {
	d := PostsController{BaseController{}}
	offset, pageSize, err := d.CalculatePagination(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	postsDto, err := reuqest.GetJsonToObj[dto.PostsDto](c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
		return
	}
	pageResp := service.PagePost(postsDto, offset, pageSize)
	response.SuccessData(c, pageResp)
}

func (PostsController) NewPagePost(c *gin.Context) {
	postsDto, _ := reuqest.GetJsonToObj[dto.PostsDto](c)
	paging := &dto.Paging{Page: (postsDto.Page), PageSize: (postsDto.PageSize)}
	paging.GetPages()
	pageResp := service.NewPagePost(postsDto, paging)
	response.SuccessData(c, pageResp)
}

// @新增博客
// @Summary 新增博客
// @Tags user
// @Description AddPosts
// @Accept  json
// @Produce json
// @Param   Title     path    string     true        "Title"
// @Param   Content     path    string     true        "Content"
// @Success 200 {string} string	"ok"
// @Router /AddPosts [post]
func (PostsController) AddPosts(c *gin.Context) {
	postMap, err := reuqest.GetJson(c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
		return
	}
	if _, ok := postMap["title"]; !ok {
		response.FailMsg(c, "博客标题不能为空")
		return
	}
	if _, ok := postMap["content"]; !ok {
		response.FailMsg(c, "博客内容不能为空")
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
		response.FailMsg(c, "未认证用户不能创建博客")
		return
	}

	posts := models.Post{}
	posts.Title = postMap["title"].(string)
	posts.Content = postMap["content"].(string)
	posts.UserId = user.Id
	posts.CreateAt = time.Now()
	service.AddPosts(&posts)
}

// 单个文章的详细信息
func (PostsController) GetPosts(c *gin.Context) {
	post, err := reuqest.GetJsonToObj[models.Post](c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
		return
	}
	if post.Id == 0 {
		response.FailMsg(c, "查询博客信息id不能为空")
		return
	}
	service.GetPost(&post)
	response.SuccessData(c, post)
}

// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func (PostsController) UpdatePost(c *gin.Context) {
	post, err := reuqest.GetJsonToObj[models.Post](c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		return
	}
	user := userInterface.(*models.User)
	service.GteUserInfo(user)

	dbPost := post
	service.GetPost(&dbPost)

	if user.Id != dbPost.UserId {
		response.FailMsg(c, "只有文章的作者才能更新自己的文章")
		return
	}
	service.UpdatePosts(&post)
	response.SuccessMsg(c, "更新成功")
}

// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func (PostsController) DelPost(c *gin.Context) {
	post, err := reuqest.GetJsonToObj[models.Post](c)
	if err != nil {
		response.FailMsg(c, "传递数据不是json")
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		return
	}
	user := userInterface.(*models.User)
	service.GteUserInfo(user)

	dbPost := post
	service.GetPost(&dbPost)

	if user.Id != dbPost.UserId {
		response.FailMsg(c, "只有文章的作者才能删除自己的文章")
		return
	}
	service.Del(&post)
	response.SuccessMsg(c, "删除成功")
}
