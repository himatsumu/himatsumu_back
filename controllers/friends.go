package controllers

import (
	"app/services"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// リクエストボディの構造体
type RequestBody struct {
    SenderId   string `json:"SenderId"`    
	ReceiverId string `json:"ReceiverId"`
}

// リクエストをDBに登録
func SendRequest(ctx echo.Context) error {
	var body RequestBody

	err := ctx.Bind(&body)
	log.Println(body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		}) 
	}

	// サービスを呼び出す
	result := services.SendRequest(body.SenderId,body.ReceiverId)

	// エラー処理
	if result.Status != 201 {
		return ctx.JSON(result.Status, result)
	}

	return ctx.JSON(result.Status, result)
}

//idを元にユーザーの情報を返す
func GetRequest(ctx echo.Context) error {
	id := ctx.Param("userId")

	// サービスを呼び出す
	result := services.GetRequest(id)

	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(result.Status,result)
	}

	return ctx.JSON(http.StatusOK, result)
}

// リクエストボディの構造体
type RegisterBody struct {
	RequestId  string `json:"RequestId"`
    SenderId   string `json:"SenderId"`    
	ReceiverId string `json:"ReceiverId"`
}


func RegisterFriend(ctx echo.Context) error {
	var body RegisterBody

	err := ctx.Bind(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		}) 
	}

	// サービスを呼び出す
	result := services.FriendRecord(body.RequestId,body.SenderId,body.ReceiverId)

	return ctx.JSON(result.Status, result)
}

func GetFriends(ctx echo.Context) error {
	id := ctx.Param("userId")

	// サービスを呼び出す
	result := services.GetFriendsByUuid(id)
	
	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(result.Status,result)
	}

	return ctx.JSON(http.StatusOK, result)
}
