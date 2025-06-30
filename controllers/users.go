package controllers

import (
	"app/services"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsersByName(ctx echo.Context) error {
	name := ctx.Param("name")

	// エラー処理
	if name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": services.UserNotFound,
		})
	}

	// サービスを呼び出す
	result := services.GetUsers(name)

	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": result.Message,
		})
	}

	return ctx.JSON(http.StatusOK, result.Data)
}
