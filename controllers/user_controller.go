package controllers

import (
	"app/services"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func CheckUser(ctx echo.Context) error {
	// リクエストからクエリパラメータを取得
	uuid := ctx.Get("user_uuid").(string)
	// UUIDが空の場合エラーを返す
	if uuid == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "UUID is required"})
	}

	// Service層の関数を呼び出す
	result, err := services.CheckUser(uuid)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check user"})
	}

	// ユーザー情報を返す
	return ctx.JSON(http.StatusOK, result)
}
