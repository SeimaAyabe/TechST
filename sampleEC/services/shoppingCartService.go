package services

import (

	// エンティティアクセル用モジュール
	"fmt"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
)

// JudgeShoppingCart は 商品データが被らないように挿入させる為の処理
func JudgeShoppingCart() {
	// 買い物カゴ一覧を取得する
	getShoppingCarts := dao.SelectAllInShoppingCart()

	fmt.Println(getShoppingCarts)

}
