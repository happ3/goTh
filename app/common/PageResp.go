package common

type PageResp struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"totalPages"`
}

func (PageResp) PageResult(offset, pageSize int, total int64, data interface{}) PageResp {

	return PageResp{
		Page:       offset + 1,
		Total:      total,
		Data:       data,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}
}
