package controller

import (
	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"
	// DBアクセス用モジュール
)

// CreateAccount は 新規会員登録を行う
func CreateAccount(c *gin.Context) {
	// 「顧客」エンティティを変数「form」として定義
	var form entity.Customer

	// バリデーション処理
	if err := c.Bind(&form); err != nil {
		// 「新規会員登録」画面に戻り、エラーメッセージを表示する
		c.HTML(http.StatusBadRequest, "signUp.html", gin.H{"err": "必須項目を入力してください"})

		// 処理を強制終了させ、後に続く処理をスキップする
		c.Abort()
	} else {
		// 入力フォーム(新規顧客情報)をサーバー側で受け取る
		//userName := c.PostForm("username")
		//mailAddress := c.PostForm("mailaddress")
		// password := c.PostForm("password")

		// 登録ユーザーが重複していた場合、「新規会員登録」画面に戻り、
		// エラーメッセージを表示する
		/*if err := createUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "signUp.html", gin.H{"err": err})
		}
		c.Redirect(302, "/")*/
	}

}
