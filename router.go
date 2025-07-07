package main

import (
	"app/controllers"
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

	authGroup := server.Group("/auth", jwtMiddleware)
	{
		authGroup.GET("/", services.GetAuthenticatedData) // http://localhost:8888/auth/

		userGroup := authGroup.Group("/user")
		{
			userGroup.GET("/", controllers.CheckUser) // http://localhost:8888/auth/user/

			userGroup.POST("/signup", controllers.Signup) // http://localhost:8888/auth/user/signup/
		}

		costumeGroup := authGroup.Group("/costume")
		{
			costumeGroup.GET("/:cos_uuid", controllers.GetCostume) // http://localhost:8888/auth/costume/
		}
	}
	return server
}
