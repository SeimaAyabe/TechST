package services

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	dao "github.com/username/sampleEC/models/dao"
	entity "github.com/username/sampleEC/models/entity"
)

// IsLoginCustomerDataExist は 入力された情報が、「顧客」テーブルのデータ内に存在するかをチェックする処理
func IsLoginCustomerDataExist(c *gin.Context, mailAddress, password string) {
	// 入力された情報が、「顧客」テーブルのデータ内に存在するかをチェックする
	getCustomer := dao.SelectCustomerByLoginInfo(mailAddress, password)

	if len(getCustomer) == 1 {
		fmt.Println(getCustomer[0].MailAddress)
		setSessionInfo(c, getCustomer)
	}
}

// setSessionInfo は セッション情報として、データをセットする処理
func setSessionInfo(c *gin.Context, customer []entity.Customer) {
	// セッションモジュールをセットする
	session := sessions.Default(c)

	// セッションにデータを保存する
	session.Set("alive", true)
	session.Set("userId", customer[0].ID)
	session.Set("userName", customer[0].UserName)

	// セッション情報をコミットする
	session.Save()
}

// DeleteSessionInfo は セッションを削除する処理
func DeleteSessionInfo(c *gin.Context) {
	// セッションモジュールをセットする
	session := sessions.Default(c)

	// セッションを削除する
	session.Clear()
	session.Save()
}
