package utils

import (
	"go-server-starter/internal/constant"
	"go-server-starter/internal/dto"
)

func NormalizePageAndPageSize(page int, pageSize int) (int, int) {
	var maxPageSize = constant.MAX_PAGE_SIZE
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = constant.DEFAULT_PAGE_SIZE
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	return page, pageSize
}

func AssemblePaginationResDto[T any](data []T, total int64, page int, pageSize int) *dto.PaginationResDto[[]T] {
	page, pageSize = NormalizePageAndPageSize(page, pageSize)
	return &dto.PaginationResDto[[]T]{
		Data:      data,
		Total:     total,
		TotalPage: (total + int64(pageSize) - 1) / int64(pageSize),
		HasNext:   total > int64(page)*int64(pageSize),
		PaginationReqDto: dto.PaginationReqDto{
			Page:     page,
			PageSize: pageSize,
		},
	}
}
