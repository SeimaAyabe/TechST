package entity

import (
	// GORM (Go言語のORM)
	"github.com/jinzhu/gorm"
)

// Product は「商品」テーブルのモデル
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;not null"`
	Price         int    `gorm:"not null"`
	// CompiledPrice string
}
