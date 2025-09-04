package controllers

import (
	"app/services"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// リクエストボディの構造体
type RegisterFolderBody struct {
	FriendUUId  string `json:"FriendUUId"`
	Date  		string `json:"Date"`
}

//フォルダーを作成
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

type Data struct {
	FriendUUId string `json:"FriendUUId"`
	Date        string `json:"Date"`
}

//写真を保存
func UplodImg(ctx echo.Context) error {
	// ファイルを取得
    file, err := ctx.FormFile("image")
    if err != nil {
		log.Println(err)
        return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		})
    }
	var body Data

	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		})
	}

	// サービスを呼び出す
	result := services.UplordImg(body.FriendUUId,body.Date,file)

	return ctx.JSON(result.Status, result)
}

type AlbumBody struct {
    FriendUUID string `json:"friend_uuid"`
}


//アルバム取得
func GetAlbums(ctx echo.Context) error {
	var body AlbumBody

	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"status":  http.StatusBadRequest,
			"message": services.InvalidRequestFormat,
		})
	}

	result := services.GetAlbums(body.FriendUUID)
	// エラー処理
	if result.Status != 200 {
		return ctx.JSON(result.Status, result)
	}

	// 正常に取得できた場合は、200ステータスと共にファイル情報をJSON形式で返す
	return ctx.JSON(http.StatusOK, result)
}