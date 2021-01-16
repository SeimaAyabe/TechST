package entity

// ShoppingCartResult は「買い物カゴ」画面を表示させる際に扱う構造体
type ShoppingCartResult struct {
	ID       int
	Name     string
	Price    int
	Quantity int
	Subtotal int
}
