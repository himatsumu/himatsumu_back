package main

import (
	"app/middleware"
	"app/services"
	"app/controllers"
	"app/utils"
	"encoding/json"
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
	authGroup.GET("/", services.GetAuthenticatedData)	// http://localhost:8888/auth/
	

	return server
}
