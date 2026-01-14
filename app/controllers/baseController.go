package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTh/app/dto"
	"goTh/app/models/reuqest"
)

type BaseController struct{}

func (BaseController) CalculatePagination(c *gin.Context) (int, int, error) {
	pageDto, err := reuqest.GetJsonToObj[dto.PageDto](c)
	if err != nil {
		return 0, 0, fmt.Errorf("未读取到分页参数")
	}
	page := pageDto.Page
	pageSize := pageDto.PageSize
	// 设置默认值
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	return offset, pageSize, nil
}
