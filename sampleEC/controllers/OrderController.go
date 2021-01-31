package controller

import (
	// 文字列と基本データ型の変換パッケージ

	// Gin

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	dbcommonlogic "github.com/username/sampleEC/models/dbcommonlogic"
	entity "github.com/username/sampleEC/models/entity"

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

	// 買い物カゴに入っている商品一覧を取得する
	getProducts := dao.SelectProductInShoppingCart(getProductID)

	// 合計金額を取得する
	productTotal := service.GetProductTotal(getProducts)

	//　「商品検索結果」画面のHTMLを返す
	c.HTML(200, "shopping-cart.html", gin.H{"getProducts": getProducts, "productTotal": productTotal})
}

// CashRegister は 「レジ」画面へ遷移する際の処理
func CashRegister(c *gin.Context) {
	// DBに接続する
	db := dbcommonlogic.Open()

	// SessionInfo型(構造体)の変数 「LoginInfo」 を定義
	var LoginInfo entity.SessionInfo

	// セッションモジュールをセットする
	session := sessions.Default(c)

	// セッションから「ユーザID」を取得する
	LoginInfo.UserID = session.Get("userId")

	// 買い物カゴ内の商品情報を取得する
	products := dao.SelectProductInShoppingCartAgain(db)

	// 合計金額を取得する
	productTotal := service.GetProductTotal(products)

	// お届け先情報を取得する
	deliveryDestination := dao.GetDeliveryDestination(LoginInfo.UserID, db)

	// DBを切断する
	db.Close()

	//　「商品検索結果」画面のHTMLを返す
	c.HTML(200, "cash-register.html", gin.H{"products": products, "productTotal": productTotal, "deliveryDestination": deliveryDestination})

}

// OrderComplete は 「注文完了」画面へ遷移する際の処理
func OrderComplete(c *gin.Context) {
	// 買い物カゴから商品を削除する
	dao.DeleteProductInShoppingCart()

	//　「注文完了」画面のHTMLを返す
	c.HTML(200, "order-complete.html", gin.H{})
}
