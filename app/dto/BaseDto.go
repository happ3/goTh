package dto

type BaseDto struct {
	Page     int `json:"page" form:"page"`         //当前页
	PageSize int `json:"pagesize" form:"pagesize"` //每页条数
}
