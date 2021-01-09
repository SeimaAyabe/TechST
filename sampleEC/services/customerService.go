package services

import (
	"fmt"

	dao "github.com/username/sampleEC/models/dao"
)

// IsLoginCustomerDataExist は 入力された情報が、「顧客」テーブルのデータ内に存在するかをチェックする処理
func IsLoginCustomerDataExist(mailAddress, password string) {
	// 入力された情報が、「顧客」テーブルのデータ内に存在するかをチェックする
	getCustomer := dao.SelectCustomer(mailAddress, password)

	fmt.Println(getCustomer[0].MailAddress)
}
