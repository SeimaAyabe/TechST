package dao

import (
	// DB共通処理
	dbcommonlogic "github.com/username/sampleEC/models/dbcommonlogic"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"
)

// InsertCustomer は 「顧客」テーブルに顧客情報を追加する
func InsertCustomer(userName, mailAddress, password string) []error {
	// DBに接続する
	db := dbcommonlogic.Open()

	// 「顧客」テーブルにレコードを入れるため、構造体を定義する
	var customer = entity.Customer{
		UserName:    userName,
		MailAddress: mailAddress,
		Password:    password,
	}

	// 「買い物カゴ」テーブルに商品情報を追加する
	db.Create(&customer)

	// DBを切断する　(return時に実行)
	defer db.Close()

	return nil

}
