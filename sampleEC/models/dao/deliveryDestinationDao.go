package dao

import (
	// 型変換パッケージ

	// DB共通処理

	"github.com/jinzhu/gorm"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"
)

// GetDeliveryDestination は お届け先情報を取得する
func GetDeliveryDestination(userID interface{}, db *gorm.DB) []entity.DeliveryDestination {
	// お届け先情報を取得するため、構造体を定義する
	var results = []entity.DeliveryDestination{}

	// お届け先情報を取得する
	db.Table("delivery_destination").Select("delivery_destination.zip_code, delivery_destination.address").Where("user_id = ? AND default_flag = ?", userID, 1).Scan(&results)

	return results

}
