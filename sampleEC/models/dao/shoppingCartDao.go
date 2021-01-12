package dao

import (
	// 型変換パッケージ
	"fmt"
	"strconv"

	// DB共通処理
	"github.com/jinzhu/gorm"
	dbcommonlogic "github.com/username/sampleEC/models/dbcommonlogic"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"
)

// InsertProductToShoppingCart は 「買い物カゴ」テーブルにレコードを追加する
func InsertProductToShoppingCart(productID string, quantity []string) {
	// DBに接続する
	db := dbcommonlogic.Open()

	// 文字列を数値に変換する
	newProductID, _ := strconv.Atoi(productID)
	compiledQuantity, _ := strconv.Atoi(quantity[0])

	// 「買い物カゴ」テーブルにレコードを入れるため、構造体を定義する
	var shoppingCart = entity.ShoppingCart{
		ShoppingCartID: 0,
		ProductID:      newProductID,
		Quantity:       compiledQuantity,
	}

	// 「買い物カゴ」テーブルに商品情報を追加する
	db.Select("ShoppingCartID", "ProductID", "Quantity").Create(&shoppingCart)

	// DBを切断する
	defer db.Close()

}

// SelectProductIDInShoppingCart は 該当する「商品ID」に合致するレコードを取得する
func SelectProductIDInShoppingCart(productID string) []entity.ShoppingCart {
	// DBに接続する
	db := dbcommonlogic.Open()

	// 買い物カゴの情報を取得するため、構造体を定義する
	var shoppingCart = []entity.ShoppingCart{}

	// DB内で検索キーワードの部分一致検索を行う
	db.Where("product_id = ?", productID).Find(&shoppingCart)

	// DBを切断する　(return時に実行)
	defer db.Close()

	// 戻り値 「買い物カゴ」内のレコード
	return shoppingCart

}

// UpdateQuantityInShoppingCart は 買い物カゴ内の数量を更新する
func UpdateQuantityInShoppingCart(productID string, addedQuantity int) {
	// DBに接続する
	db := dbcommonlogic.Open()

	// 買い物カゴの情報を取得するため、構造体を定義する
	var shoppingCart = []entity.ShoppingCart{}

	// 買い物カゴ一覧を取得する
	db.Model(&shoppingCart).Where("product_id = ?", productID).Update("quantity", addedQuantity)

	// DBを切断する
	defer db.Close()
}

// SelectProductInShoppingCart は 買い物カゴに入っている商品一覧を取得する
func SelectProductInShoppingCart() ([]entity.ShoppingCartResult, int) {
	// DBに接続する
	db := dbcommonlogic.Open()

	// 買い物カゴ内の商品情報を取得するため、構造体を定義する
	var results = []entity.ShoppingCartResult{}

	// 買い物カゴに追加された商品情報の一覧を取得する
	db.Table("product").Select("product.id, product.name, product.price, shopping_cart.quantity").Joins("inner join shopping_cart on product_id = product.id").Scan(&results)

	fmt.Println(results)

	// 小計を算出し、「買い物カゴ」テーブルに挿入する
	subtotal := InsertSubtotalToShoppingCart(db, results)

	// DBを切断する　(return時に実行)
	defer db.Close()

	// 戻り値 「買い物カゴ内の商品情報一覧」
	return results, subtotal
}

// InsertSubtotalToShoppingCart は 小計を算出し、「買い物カゴ」テーブルに挿入する
func InsertSubtotalToShoppingCart(db *gorm.DB, results []entity.ShoppingCartResult) int {
	// 小計を算出
	subtotal := results[0].Price * results[0].Quantity

	// 「買い物カゴ」テーブルの構造体を定義する
	var shoppingCart = []entity.ShoppingCart{}

	// 「買い物カゴ」テーブル内の「小計」を更新する
	db.Model(&shoppingCart).Where("product_id = ?", results[0].ID).Update("subtotal", subtotal)

	fmt.Println(subtotal)

	// 戻り値 「小計」
	return subtotal
}
