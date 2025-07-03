package controllers

import (
	"app/custom_error"
	"app/services"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func CheckUser(ctx echo.Context) error {
	// リクエストからクエリパラメータを取得
	uuid := ctx.Get("user_uuid").(string)
	// UUIDが空の場合エラーを返す
	if uuid == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "UUIDが空です",
		})
	}

	// Service層の関数を呼び出す
	result, err := services.CheckUser(uuid)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "予期しないエラーが発生しました",
		})
	}

	// ユーザー情報を返す
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   result,
	})
}

func Signup(ctx echo.Context) error {
	// リクエストを構造体にバインド
	req := new(services.SignupRequest) // サービス層で定義した型を使う
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "リクエストの中身が不正です",
		})
	}

	// Service層の関数を呼び出す
	result, err := services.Signup(req, ctx)
	if err != nil {
		// カスタムエラーかチェック
		if customErr, ok := err.(*custom_error.CustomError); ok {
			return ctx.JSON(customErr.Code, map[string]interface{}{
				"status": customErr.Code,
				"message":  customErr.Message,
			})
		}

		// 予期しないエラーの場合
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusInternalServerError,
			"message":  "予期しないエラーが発生しました",
		})
	}

	// ユーザー情報を返す
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "ユーザーを作成しました",
		"user_id": result,
	})
}
