package controller

import (
	// 文字列と基本データ型の変換パッケージ

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"
)

// FetchAllProducts は 全ての商品情報を取得する
func FetchAllProducts(c *gin.Context) {
	resultProducts := dao.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}
