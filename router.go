package main

import (
	"app/middleware"
	"app/services"
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

	// ユーザーサービス
	userService := services.NewUserService()
	server.GET("/auth", userService.GetAuthenticatedData, jwtMiddleware)

	return server
}
