package dao

import (
	// 型変換パッケージ
	"strconv"

	// DB共通処理
	dbcommon "github.com/username/sampleEC/models/dbcommon"
	entity "github.com/username/sampleEC/models/entity"
)

// InsertProductToShoppingCart は 「カゴ」テーブルに商品を追加する
func InsertProductToShoppingCart(productID string) {
	// DBに接続する
	db := dbcommon.Open()

	newProductID, _ := strconv.Atoi(productID)

	// insert
	var shoppingCart = entity.ShoppingCart{
		ShoppingCartID: 0,
		ProductID:      newProductID,
	}

	db.Select("ShoppingCartID", "ProductID").Create(&shoppingCart)

	defer db.Close()
}
