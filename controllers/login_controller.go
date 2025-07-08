package controllers

import (
	"app/services"
	"net/http"

	"github.com/labstack/echo"
)

func CheckUser(ctx echo.Context) error {
	// リクエストからクエリパラメータを取得
	uuid := ctx.Get("user_uuid").(string)

	// Service層の関数を呼び出す
	result := services.CheckUser(uuid)
	if result.Status != http.StatusOK {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": result.Message,
		})
	}

	// ユーザー情報を返す
	return ctx.JSON(http.StatusOK, result)
}

func Signup(ctx echo.Context) error {
	// リクエストを構造体にバインド
	req := new(services.SignupRequest) // サービス層で定義した型を使う
	ctx.Bind(req)
	// リクエストからユーザー情報を取得
	user_uuid := ctx.Get("user_uuid").(string)

	// Service層の関数を呼び出す
	result := services.Signup(req, user_uuid)

	// ユーザー情報を返す
	return ctx.JSON(http.StatusCreated, result)
}
