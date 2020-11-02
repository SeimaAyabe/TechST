package controller

import (
	// 文字列と基本データ型の変換パッケージ

	// Gin
	"github.com/gin-gonic/gin"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
)

// AddToShoppingCart は 買い物カゴへ商品を追加する
func AddToShoppingCart(c *gin.Context) {
	// 「商品」テーブルの「ID」をサーバー側で受け取る
	getProductID := c.Param("ID")

	// 「買い物カゴ」テーブルに、商品IDを追加する
	dao.InsertProductToShoppingCart(getProductID)

	//　「商品詳細」画面のHTMLを返す
	c.HTML(200, "top.html", gin.H{})
}
