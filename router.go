package main

import (
	"net/http"
	"github.com/labstack/echo"
)
func InitServer() *echo.Echo {
	// サーバー作成
	server := echo.New()


	server.POST("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})
	
	return server
}