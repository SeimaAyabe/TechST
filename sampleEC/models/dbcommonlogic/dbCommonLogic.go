package dbcommon

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Open DBに接続するメソッド
func Open() *gorm.DB {
	// DBに関する情報を定義
	DBMS := "mysql"
	USER := "root"
	PASS := "thamen1451"
	PROTOCOL := "tcp(SampleEC-db:3306)"
	DBNAME := "thamen1451"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	// DB接続エラー時の処理
	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// DBが接続されたことを示すログを出力
	fmt.Println("db connected: ", &db)

	return db
}
