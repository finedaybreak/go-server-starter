package dto

type PaginationReqDto struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type PaginationResDto[T any] struct {
	Data T `json:"data"`
	PaginationReqDto
	Total     int64 `json:"total"`
	TotalPage int64 `json:"totalPage"`
	HasNext   bool  `json:"hasNext"`
}
