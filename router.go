package main

import (
	"app/controllers"
	"app/middleware"
	"app/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func InitServer() *echo.Echo {
	// サーバー作成
	server := echo.New()

	// CORSミドルウェア
	server.Use(middleware.CORSMiddleware())

	publicKey, err := utils.LoadPublicKey("keys/public.key")
	if err != nil {
		log.Fatal(err)
	}

	// JWT認証ミドルウェア
	jwtMiddleware := middleware.JWTAuthMiddleware(publicKey)

	server.POST("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World! from go-server.")
	})

	authGroup := server.Group("/auth", jwtMiddleware)
	{

		//ユーザー
		userGroup := authGroup.Group("/user")
		{
			userGroup.GET("/", controllers.CheckUser) // http://localhost:8888/auth/user/

			userGroup.POST("/signup", controllers.Signup) // http://localhost:8888/auth/user/signup/

			userGroup.GET("/:userId", controllers.GetUsersById) // http://localhost:8888/auth/user/:userId

		}

		//リクエスト
		requestGroup := authGroup.Group("/request")
		{
			//フレンド申請送信
			requestGroup.POST("/send/:receiverUuid",controllers.SendRequest) // http://localhost:8888/auth/request/send/:receiverUuid
			//フレンド申請受信
			requestGroup.GET("/:userId",controllers.GetRequest) // http://localhost:8888/auth/request/:userId
			//フレンド登録
			requestGroup.POST("/register",controllers.RegisterFriend)// http://localhost:8888/auth/request/register
		}

		//フレンド
		friendGroup := authGroup.Group("/friend")
		{
			//フレンド一覧検索
			friendGroup.GET("/",controllers.GetFriends)// http://localhost:8888/auth/friend/
		}

		//キャラクター
		characterGroup := authGroup.Group("/character")
		{
			_= characterGroup
		}

		//クエスト
		questGroup := authGroup.Group("/quest")
		{
			questGroup.POST("/quests", controllers.GenerateQuests)	// http://localhost:8888/auth/quest/quests

			questGroup.POST("/create", controllers.CreateQuest)	// http://localhost:8888/auth/quest/create

			questGroup.POST("/check", controllers.CheckQuest) // http://localhost:8888/auth/quest/check
		}

		//アルバム
		albumGroup := authGroup.Group("/album")
		{
			_=albumGroup
		}



	}

	return server
}
