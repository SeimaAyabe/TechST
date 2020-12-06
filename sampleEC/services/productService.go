package services

import (
	"fmt"
	"strconv"

	// エンティティアクセル用モジュール
	entity "github.com/username/sampleEC/models/entity"
)

// CompilePrice は 商品の値段を表示させる為の処理
func CompilePrice(productList []entity.Product) {

	fmt.Println(productList[0].Price)

	stringPrice := strconv.Itoa(productList[0].Price)

	fmt.Println(len(stringPrice))

	var commaPrice string

	var priceCount int

	for i := len(stringPrice); i > 0; i-- {

		commaPrice += stringPrice[i-1 : i]
		priceCount++
		if priceCount%3 == 0 && priceCount != len(stringPrice) {
			commaPrice += ","
		}
		fmt.Println(commaPrice)
	}

	// 文字列を逆順に並び替える
	productList[0].CompiledPrice = ReverseString(commaPrice)

	fmt.Println(productList)

}

// ReverseString は文字列を逆さにする
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
