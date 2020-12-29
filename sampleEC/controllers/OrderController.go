package controller

import (
	// 文字列と基本データ型の変換パッケージ

	// Gin
	"fmt"

	"github.com/gin-gonic/gin"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
)

// AddToShoppingCart は 買い物カゴへ商品を追加する
func AddToShoppingCart(c *gin.Context) {
	// 「商品」テーブルの「ID」と、商品の数量をサーバー側で受け取る
	c.Request.ParseForm()
	getProductID := c.Param("ID")
	getQuantity := c.Request.Form["cnt"]

	fmt.Println(getQuantity)

	// 「買い物カゴ」テーブルに、商品IDを追加する
	dao.InsertProductToShoppingCart(getProductID)

	// 買い物カゴに入っている商品一覧を取得する
	getProducts := dao.SelectProductInShoppingCart()

	fmt.Println(getProducts)

	//　「商品検索結果」画面のHTMLを返す
	c.HTML(200, "shopping-cart.html", gin.H{"getProducts": getProducts})
}
