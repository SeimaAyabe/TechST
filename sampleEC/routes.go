package main

import (
	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// セッションのパッケージ
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラー
	controller "github.com/username/sampleEC/controllers"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// 静的ファイルのパスを指定
	router.Static("/views", "./views")

	// HTMLファイルのパスを指定
	router.LoadHTMLGlob("./views/html/*/**.html")

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("sampleECSession", store))

	// 「トップ」画面へのリクエストに対するアクション
	router.GET("/", controller.SessionCheck)

	// 「商品検索結果」画面へのリクエストに対するアクション
	router.GET("/SearchResult", controller.SearchProducts)

	// 「商品詳細」画面へのリクエストに対するアクション
	router.GET("/ProductDetail/:ID", controller.ProductDetail)

	// 「買い物カゴ」画面へのリクエストに対するアクション
	router.POST("/ShoppingCart/:ID", controller.AddToShoppingCart)

	// 「レジ」画面へのリクエストに対するアクション
	router.GET("/CashRegister", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "cash-register.html", gin.H{})
	})

	// 「新規会員登録」画面へのリクエストに対するアクション
	router.GET("/SignUp", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create-account.html", gin.H{})
	})

	// 「会員登録」ボタン押下後のリクエストに対するアクション
	router.POST("/CreateAccount", controller.CreateAccount)

	// 「ログイン」画面へのリクエストに対するアクション
	router.GET("/Login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// 「ログイン」ボタン(「ログイン」画面)押下後のリクエストに対するアクション
	router.POST("/Signin", controller.Signin)

	// 「ログアウト」ボタン(「トップ」画面)押下後のリクエストに対するアクション
	router.GET("/Logout", controller.Logout)

	// 「マイページ」画面へのリクエストに対するアクション
	router.GET("/Mypage", controller.GetAccountInfo)

	// 「お届け先新規作成」画面へのリクエストに対するアクション
	router.GET("/CreateDeliveryDestination", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create-delivery-destination.html", gin.H{})
	})

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	router.StaticFS("/shoppingapp", http.Dir("./views/vuetify"))

	// 全ての商品情報のJSONを返す
	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	// １つの商品情報の状態のJSONを返す
	router.GET("/fetchProduct", controller.FindProduct)

	// 商品情報をDBへ登録する
	// router.POST("/addProduct", controller.AddProduct)

	// 商品情報の状態を変更する
	router.POST("/changeStateProduct", controller.ChangeStateProduct)

	// 商品情報を削除する
	router.POST("/deleteProduct", controller.DeleteProduct)

	if err := router.Run(":8081"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
