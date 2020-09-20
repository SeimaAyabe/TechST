package dao

// フォーマットI/O
// エンティティ (データベースのテーブルの行に対応)

// DB接続する
/*func openCustomerTable() *gorm.DB {
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
	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connected: ", &db)
	return db
}

// FindAllProducts は 商品テーブルのレコードを全件取得する
func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := open()
	// select
	db.Order("ID asc").Find(&products)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return products
}*/
