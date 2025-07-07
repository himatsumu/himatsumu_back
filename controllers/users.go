package controllers

import (
	"app/services"
	"net/http"

	"github.com/labstack/echo"
)

//名前を元にユーザーの情報を返す
func GetUsersByName(ctx echo.Context) error {
	name := ctx.Param("name")

	// サービスを呼び出す
	result := services.GetUsersByName(name)

	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(result.Status,result)
	}

	return ctx.JSON(http.StatusOK, result)
}


//idを元にユーザーの情報を返す
func GetUsersById(ctx echo.Context) error {
	id := ctx.Param("userId")

	// サービスを呼び出す
	result := services.GetUsersById(id)

	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(result.Status, result)
	}
	return ctx.JSON(http.StatusOK, result)
}
