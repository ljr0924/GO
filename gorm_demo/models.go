package gorm_demo

import (
	"gorm.io/gorm"
	"strings"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// AfterFind 钩子函数，查询后调用
func (p *Product) AfterFind(tx *gorm.DB) (err error) {
	codeLen := len(p.Code)
	if codeLen > 1 {
		p.Code = p.Code[0:1] + strings.Repeat("*", codeLen-1)
	}
	return
}
