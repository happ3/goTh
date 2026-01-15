package dto

import "math"

type Paging struct {
	Page      int `json:"page" form:"page"`           //当前页
	PageSize  int `json:"pagesize" form:"pagesize"`   //每页条数
	Total     int `json:"total" form:"total"`         //总条数
	PageCount int `json:"pagecount" form:"pagecount"` //总页数
	Offset    int `json:"offset" form:"offset"`       //起始条数
}

// 获取分页信息
func (p *Paging) GetPages() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	page_count := math.Ceil(float64(p.Total) / float64(p.PageSize))
	p.Offset = p.PageSize * (p.Page - 1)
	p.PageCount = int(page_count)
}
