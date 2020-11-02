package dao

import ( // フォーマットI/O
	// GORM (Go言語のORM)
	// DB共通処理
	dbcommonlogic "github.com/username/sampleEC/models/dbcommonlogic"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"
)

// FindAllProducts は 商品テーブルのレコードを全件取得する
func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := dbcommonlogic.Open()
	// select
	db.Order("ID asc").Find(&products)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return products
}

// SearchProducts は 検索キーワードに該当する「商品」テーブルのレコードを取得する
func SearchProducts(searchedProducts string) []entity.Product {
	// Product型の変数 「product」 を定義
	product := []entity.Product{}

	// DB接続
	db := dbcommonlogic.Open()

	// DB内で検索キーワードの部分一致検索を行う
	db.Where("name LIKE ?", "%"+searchedProducts+"%").Find(&product)

	// DB切断 (return時に実行)
	defer db.Close()

	// 戻り値 「商品」エンティティ
	return product
}

// ProductDetail は 選択した「商品」テーブルのレコードを取得する
func ProductDetail(getProduct string) []entity.Product {
	// Product型の変数 「product」 を定義
	product := []entity.Product{}

	// DB接続
	db := dbcommonlogic.Open()

	// DB内で検索キーワードの部分一致検索を行う
	db.Where("id = ?", getProduct).Find(&product)

	// DB切断 (return時に実行)
	defer db.Close()

	// 戻り値 「商品」エンティティ
	return product
}

// FindProduct は 商品テーブルのレコードを１件取得する
func FindProduct(productID int) []entity.Product {
	product := []entity.Product{}

	db := dbcommonlogic.Open()
	// select
	db.First(&product, productID)
	defer db.Close()

	return product
}

// UpdateStateProduct は 商品テーブルの指定したレコードの状態を変更する
func UpdateStateProduct(productID int, productState int) {
	product := []entity.Product{}

	db := dbcommonlogic.Open()
	// update
	db.Model(&product).Where("ID = ?", productID).Update("State", productState)
	defer db.Close()
}

// DeleteProduct は 商品テーブルの指定したレコードを削除する
func DeleteProduct(productID int) {
	product := []entity.Product{}

	db := dbcommonlogic.Open()
	// delete
	db.Delete(&product, productID)
	defer db.Close()
}
