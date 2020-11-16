package dbcommon

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/username/sampleEC/models/entity"
)

// Open DBに接続するメソッド
func Open() *gorm.DB {
	// DBに関する情報を定義
	DBMS := "mysql"
	USER := "root"
	PASS := "thamen1451"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "thamen1451"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Customer{})

	fmt.Println("db connected: ", &db)
	return db
}
