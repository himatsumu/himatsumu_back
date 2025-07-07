package controllers

import (
	"app/services"
	"app/custom_error"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCostume(ctx echo.Context) error {

	// UUIDを取得する
	cos_uuid := ctx.Param("cos_uuid")

	// サービス層の関数を呼び出す
	result, err := services.GetCostume(cos_uuid)
	if err != nil {
		// カスタムエラーかチェック
		if customErr, ok := err.(*custom_error.CustomError); ok {
			return ctx.JSON(customErr.Code, echo.Map{
				"status": customErr.Code,
				"message":  customErr.Message,
			})
		}

		// 予期しないエラーの場合
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"message":  "予期しないエラーが発生しました",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"message": "コスチューム情報を取得しました",
		"data":   result,
	})
}