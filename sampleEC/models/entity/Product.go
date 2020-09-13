package entity

// Product は「商品」テーブルのモデル
type Product struct {
	ID    int    `gorm:"primary_key;not null"`
	Name  string `gorm:"type:varchar(200);not null"`
	Memo  string `gorm:"type:varchar(400)"`
	State int    `gorm:"not null"`
}
