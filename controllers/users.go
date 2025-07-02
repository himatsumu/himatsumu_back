package controllers

import (
	"app/services"
	"net/http"

	"github.com/labstack/echo"
)

//名前を元にユーザーの情報を返す
func GetUsersByName(ctx echo.Context) error {
	name := ctx.Param("name")

	// エラー処理
	if name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": services.UserNotFound,
		})
	}

	// サービスを呼び出す
	result := services.GetUsersByName(name)

	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": result.Message,
		})
	}

	return ctx.JSON(http.StatusOK, result.Data)
}


//idを元にユーザーの情報を返す
func GetUsersById(ctx echo.Context) error {
	id := ctx.Param("id")

	// エラー処理
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": services.UserNotFound,
		})
	}

	// サービスを呼び出す
	result := services.GetUsersById(id)

	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": result.Message,
		})
	}

	return ctx.JSON(http.StatusOK, result.Data)
}
