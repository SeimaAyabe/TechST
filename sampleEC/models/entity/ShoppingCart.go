package entity

// ShoppingCart は「買い物カゴ」テーブルのモデル
type ShoppingCart struct {
	ShoppingCartID int `gorm:"primary_key;not null"`
	ProductID      int `gorm:"not null"`
}
