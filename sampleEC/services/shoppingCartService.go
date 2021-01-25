package services

import (

	// エンティティアクセル用モジュール

	"strconv"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
	entity "github.com/username/sampleEC/models/entity"
)

// JudgeShoppingCart は 商品データが被らないように挿入させる為の処理
func JudgeShoppingCart(productID string, quantity []string) {
	// 選択した商品のIDが既に「買い物カゴ」テーブル内に存在するかどうかをチェックする
	getProductIDInShoppingCart := dao.SelectProductIDInShoppingCart(productID)

	// 文字列を数値に変換する
	compiledQuantity, _ := strconv.Atoi(quantity[0])

	// 既に存在していた場合、数量を加算して更新する
	if len(getProductIDInShoppingCart) == 1 {
		// 加算処理を行う
		addedQuantity := getProductIDInShoppingCart[0].Quantity + compiledQuantity

		// 数量の更新を行う
		dao.UpdateQuantityInShoppingCart(productID, addedQuantity)

		// 存在しない場合、新たなレコードを追加する
	} else {
		// 「買い物カゴ」テーブルに、レコードを追加する
		dao.InsertProductToShoppingCart(productID, quantity)

		// 数量の更新を行う
		dao.UpdateQuantityInShoppingCart(productID, compiledQuantity)
	}

}

// GetProductTotal は 合計金額を取得する処理
func GetProductTotal(product []entity.ShoppingCartResult) int {
	// 数値型で「合計金額」を定義する
	var productTotal int

	// 買い物カゴに入っている商品の数分、処理を繰り返す
	for i := 0; i < len(product); i++ {
		productTotal += product[i].Subtotal
	}

	// 戻り値 「合計金額」
	return productTotal
}
