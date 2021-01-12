package controller

import (
	// 文字列と基本データ型の変換パッケージ

	// Gin

	"github.com/gin-gonic/gin"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"

	// サービスアクセス用モジュール
	service "github.com/username/sampleEC/services"
)

// AddToShoppingCart は 買い物カゴへ商品を追加する
func AddToShoppingCart(c *gin.Context) {
	// 「商品」テーブルの「ID」と、商品の数量をサーバー側で受け取る
	c.Request.ParseForm()
	getProductID := c.Param("ID")
	getQuantity := c.Request.Form["cnt"]

	// 「買い物カゴ」テーブル内に、既に該当の商品が存在する場合、数量を加算する為の処理
	service.JudgeShoppingCart(getProductID, getQuantity)

	// 買い物カゴに入っている商品一覧、小計を取得する
	getProducts, subtotal := dao.SelectProductInShoppingCart()

	// 小計を表示させる為の処理
	// service.DisplaySubtotal()

	//　「商品検索結果」画面のHTMLを返す
	c.HTML(200, "shopping-cart.html", gin.H{"getProducts": getProducts, "subtotal": subtotal})
}
