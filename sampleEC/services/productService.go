package services

import (
	"strconv"

	// エンティティアクセル用モジュール
	entity "github.com/username/sampleEC/models/entity"
)

// CompilePrice は 商品の値段を表示させる為の処理
func CompilePrice(productList []entity.Product) {
	// 数値を文字列に変換する
	stringPrice := strconv.Itoa(productList[0].Price)

	// カンマ入りの値段を表示させる為の文字列を定義
	var commaPrice string

	// カンマの入れる位置を判断する為の数値を定義
	var priceCount int

	// 文字列の長さ分、処理を繰り返す
	for i := len(stringPrice); i > 0; i-- {
		// 文字列の後ろから文字を代入していく
		commaPrice += stringPrice[i-1 : i]

		// 数値をインクリメント
		priceCount++

		// カンマを入れる桁かどうかを判定
		if priceCount%3 == 0 && priceCount != len(stringPrice) {
			// カンマを挿入
			commaPrice += ","
		}
	}

	// 文字列を逆順に並び替える
	productList[0].CompiledPrice = ReverseString(commaPrice)

}

// ReverseString は文字列を逆さにする
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
