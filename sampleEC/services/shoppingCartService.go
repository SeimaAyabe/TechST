package services

import (

	// エンティティアクセル用モジュール

	"fmt"
	"strconv"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
)

// JudgeShoppingCart は 商品データが被らないように挿入させる為の処理
func JudgeShoppingCart(productID string, quantity []string) {
	// 買い物カゴ一覧を取得する
	getShoppingCarts := dao.SelectAllInShoppingCart()

	// 買い物カゴ内のデータを１つずつチェックし、選択した商品が既に入っているかどうかを確認する
	for i := 0; i < len(getShoppingCarts); i++ {
		// 既に商品が入っていた場合、加算する
		if productID == strconv.Itoa(getShoppingCarts[i].ProductID) {
			// 文字列を数値に変換
			compiledQuantity, _ := strconv.Atoi(quantity[0])

			// 加算処理
			addedQuantity := getShoppingCarts[i].Quantity + compiledQuantity

			fmt.Println(addedQuantity)
		}
	}

}
