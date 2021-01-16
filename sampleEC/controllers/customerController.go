package controller

import (
	// HTTPを扱うパッケージ
	"fmt"
	"net/http"

	// Gin
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	// エンティティ (データベースのテーブルの行に対応)
	entity "github.com/username/sampleEC/models/entity"

	// DBアクセス用モジュール
	dao "github.com/username/sampleEC/models/dao"

	// サービスアクセス用モジュール
	service "github.com/username/sampleEC/services"
)

// CreateAccount は 新規会員登録を行う
func CreateAccount(c *gin.Context) {
	// 「顧客」エンティティを変数「form」として定義
	var form entity.Customer

	// バリデーション処理
	if err := c.Bind(&form); err != nil {
		// 「新規会員登録」画面に戻り、エラーメッセージを表示する
		c.HTML(http.StatusBadRequest, "create-account.html", gin.H{"err": "必須項目を入力してください"})

		// 処理を強制終了させ、後に続く処理をスキップする
		c.Abort()
	} else {
		// 入力フォーム(新規顧客情報)をサーバー側で受け取る
		userName := c.PostForm("username")
		mailAddress := c.PostForm("mailaddress")
		password := c.PostForm("password")

		// 登録ユーザーが重複していた場合、「新規会員登録」画面に戻り、
		// エラーメッセージを表示する
		if err := dao.InsertCustomer(userName, mailAddress, password); err != nil {
			c.HTML(http.StatusBadRequest, "create-account.html", gin.H{"err": "ユーザーは既に登録済です"})
		}

		// 「トップ」画面にリダイレクトする
		c.Redirect(http.StatusFound, "/")
	}

}

// Signin は ログイン処理を行う
func Signin(c *gin.Context) {
	// 「ログイン」画面にて入力された値をサーバー側で受け取る
	mailAddress := c.PostForm("mailaddress")
	password := c.PostForm("password")

	// 入力された情報が、「顧客」テーブルのデータ内に存在するかをチェックする処理
	service.IsLoginCustomerDataExist(c, mailAddress, password)

	// 「トップ」画面にリダイレクトする
	c.Redirect(http.StatusFound, "/")

}

// Logout は ログアウト処理を行う
func Logout(c *gin.Context) {
	// セッションを削除する処理
	service.DeleteSessionInfo(c)

	// 「トップ」画面にリダイレクトする
	c.Redirect(http.StatusFound, "/")

}

// SessionCheck は セッションにデータが保持されているかどうかを判定する処理
func SessionCheck(c *gin.Context) {
	// SessionInfo型(構造体)の変数 「LoginInfo」 を定義
	var LoginInfo entity.SessionInfo

	// デフォルトで "false" を設定する
	LoginInfo.IsSessionAlive = false

	// セッションモジュールをセットする
	session := sessions.Default(c)

	// セッションから「ユーザ名」を取得する
	LoginInfo.UserName = session.Get("userName")
	fmt.Println(LoginInfo.UserName)

	// セッションが保存されていた場合、セッションフラグに"true"をセットする
	if LoginInfo.UserName != nil {
		LoginInfo.IsSessionAlive = true
	}

	//　「トップ」画面のHTMLを返す
	c.HTML(200, "top.html", gin.H{"loginInfo": LoginInfo})

}

// GetAccountInfo は マイページを表示させるために必要な会員情報を取得する処理
func GetAccountInfo(c *gin.Context) {
	// SessionInfo型(構造体)の変数 「LoginInfo」 を定義
	var LoginInfo entity.SessionInfo

	// (こっちは後で消すよ！！)デフォルトで "false" を設定する
	// LoginInfo.IsSessionAlive = false

	// セッションモジュールをセットする
	session := sessions.Default(c)

	// セッションから「ユーザID」を取得する
	LoginInfo.UserID = session.Get("userId")

	// 「ユーザID」を元に、「顧客」テーブルのデータ内に存在するかをチェックする
	getCustomer := dao.SelectCustomerByUserID(LoginInfo.UserID)

	fmt.Println(getCustomer)

	//　「トップ」画面のHTMLを返す
	c.HTML(200, "mypage.html", gin.H{"customer": getCustomer})

}
