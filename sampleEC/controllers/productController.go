package controller

import (
	// 文字列と基本データ型の変換パッケージ
	strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
)

// 商品の購入状態を定義
const (
	// NotPurchased は 未購入
	NotPurchased = 0

	// Purchased は 購入済
	Purchased = 1
)

// FetchAllProducts は 全ての商品情報を取得する
func FetchAllProducts(c *gin.Context) {
	resultProducts := dao.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// SearchProducts は 検索キーワードに関連した商品情報を取得する
func SearchProducts(c *gin.Context) {
	// 検索キーワードをサーバー側で受け取る
	searchedProducts := c.Query("search")

	// モデル側で検索キーワードに該当する「商品」テーブルのレコードを取得する
	resultProducts := dao.SearchProducts(searchedProducts)

	//　「商品検索結果」画面のHTMLを返す
	c.HTML(200, "top.html", gin.H{"resultProducts": resultProducts})
}

// FindProduct は 指定したIDの商品情報を取得する
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productID, _ := strconv.Atoi(productIDStr)

	resultProduct := dao.FindProduct(productID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProduct)
}

// AddProduct は 商品をDBへ登録する
func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	var product = entity.Product{
		Name:  productName,
		Memo:  productMemo,
		State: NotPurchased,
	}

	dao.InsertProduct(&product)
}

// ChangeStateProduct は 商品情報の状態を変更する
func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	// 商品状態が未購入の場合
	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	dao.UpdateStateProduct(productID, changeState)
}

// DeleteProduct は 商品情報をDBから削除する
func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)

	dao.DeleteProduct(productID)
}
