package dao

import (
	// 型変換パッケージ
	"strconv"

	// DB共通処理
	dbcommonlogic "github.com/username/sampleEC/models/dbcommonlogic"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"
)

// InsertProductToShoppingCart は 「買い物カゴ」テーブルに商品を追加する
func InsertProductToShoppingCart(productID string) {
	// DBに接続する
	db := dbcommonlogic.Open()

	// String型をint型に変換する
	newProductID, _ := strconv.Atoi(productID)

	// 「買い物カゴ」テーブルにレコードを入れるため、構造体を定義する
	var shoppingCart = entity.ShoppingCart{
		ShoppingCartID: 0,
		ProductID:      newProductID,
	}

	// 「買い物カゴ」テーブルに商品情報を追加する
	db.Select("ShoppingCartID", "ProductID").Create(&shoppingCart)

	// DBを切断する
	defer db.Close()
}
