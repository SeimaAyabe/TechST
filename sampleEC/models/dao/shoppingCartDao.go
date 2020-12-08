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

	// DBを切断する　(return時に実行)
	defer db.Close()

}

// SelectProductInShoppingCart は 買い物カゴに入っている商品一覧を取得する
func SelectProductInShoppingCart() []entity.ShoppingCartResult {
	// DBに接続する
	db := dbcommonlogic.Open()

	// 買い物カゴ内の商品情報を取得するため、構造体を定義する
	var results = []entity.ShoppingCartResult{}

	// 買い物カゴに追加された商品情報の一覧を取得する
	db.Table("product").Select("product.name, product.price").Joins("left join shopping_cart on product_id = product.id").Scan(&results)

	// DBを切断する　(return時に実行)
	defer db.Close()

	// 戻り値 「買い物カゴ内の商品情報一覧」
	return results
}
