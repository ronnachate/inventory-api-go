package repository

import "gorm.io/gorm"

type pagination struct {
	limit int
	page  int
}

func NewPagination(page int, rows int) *pagination {
	return &pagination{limit: rows, page: page}
}

func (p *pagination) PageResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Offset(offset).
		Limit(p.limit)
}
