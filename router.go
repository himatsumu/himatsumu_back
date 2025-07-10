package main

import (
	"app/controllers"
	"app/middleware"
	"app/utils"
	"app/mocks"
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

		userGroup := authGroup.Group("/user")
		{
			userGroup.GET("/", controllers.CheckUser) // http://localhost:8888/auth/user/

			userGroup.POST("/signup", controllers.Signup) // http://localhost:8888/auth/user/signup/
		}
	}

	//モックサーバーのエンドポイント
	mockGroup := server.Group("/mock")
	{
		mockGroup.GET("/", func(ctx echo.Context) error {
			return ctx.JSON(http.StatusOK, "This is mock endpoint.")
		})

		//モックのユーザーグループ
		//mUserGroup := mockGroup.Group("/user")
		{

		}

		//モックのフレンドグループ
		mFriendGroup := mockGroup.Group("/friend")
		{
			//フレンド一覧
			mFriendGroup.GET("/", mocks.MockGetFriends)			// http://localhost:8888/mock/friend/

			//フレンド情報
			mFriendGroup.GET("/:id", mocks.MockGetFriendById)	// http://localhost:8888/mock/friend/ayaka
		}

		mQuestGroup := mockGroup.Group("/quest")
		{
			mQuestGroup.GET("/", mocks.MockGetQuests)			// http://localhost:8888/mock/quest/
		}

	}

	return server
}
