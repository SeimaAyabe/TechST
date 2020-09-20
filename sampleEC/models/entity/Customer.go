package entity

import (
	// GORM (Go言語のORM)
	"github.com/jinzhu/gorm"
)

// Customer は「顧客」テーブルのモデル
type Customer struct {
	gorm.Model
	UserName    string `form:"username" binding:"required"`
	MailAddress string `form:"mailaddress" binding:"required" gorm:"unique;not null"`
	Password    string `form:"password" binding:"required"`
}
