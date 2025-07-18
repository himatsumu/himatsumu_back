package controllers

import (
	"app/services"
	"app/models"
	"net/http"

	"github.com/labstack/echo"
)

func GenerateQuests(ctx echo.Context) error {
	// リクエストからクエリパラメータを取得
	userUuid := ctx.Get("user_uuid").(string)
	req := new(services.GenerateQuestsRequest)
	err := ctx.Bind(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		})
	}

	// Service層の関数を呼び出す
	result := services.GenerateQuests(userUuid, *req)

	// エラー処理
	if result.Status != http.StatusCreated {
		return ctx.JSON(result.Status, result)
	}

	// ユーザー情報を返す
	return ctx.JSON(http.StatusCreated, result)
}

func CreateQuest(ctx echo.Context) error {
	// リクエストからクエリパラメータを取得
	userUuid := ctx.Get("user_uuid").(string)
	req := new(models.CreateQuestRequest)
	err := ctx.Bind(req)

	// Service層の関数を呼び出す
	result := services.CreateQuest(userUuid, *req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	// エラー処理
	if result.Status != http.StatusCreated {
		return ctx.JSON(result.Status, result)
	}

	// ユーザー情報を返す
	return ctx.JSON(http.StatusCreated, result)
}