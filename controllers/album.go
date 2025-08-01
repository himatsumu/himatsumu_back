package controllers

import (
	"app/services"
	"net/http"

	"github.com/labstack/echo"
)

// リクエストボディの構造体
type RegisterFolderBody struct {
	FriendUUId  string `json:"FriendUUId"`
	Date  		string `json:"Date"`
}

func RegisterFolder(ctx echo.Context) error {
	var body RegisterFolderBody

	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		})
	}

	// サービスを呼び出す
	result := services.CreateFolder(body.FriendUUId, body.Date)

	return ctx.JSON(result.Status, result)
}
