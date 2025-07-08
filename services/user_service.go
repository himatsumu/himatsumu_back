package services

import (
	"app/models"
	"fmt"
	"time"

	"net/http"
)

// ユーザー情報を返す
func CheckUser(uuid string) Result {
	// UUIDが空か確認
	if uuid == "" {
		return Result{
			Message: EmptyUUID,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// ユーザを取得する(なければいないことを返す)
	result, err := models.CheckUser(uuid)
	if err != nil {
		return Result{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	// 結果をそのまま返す（IsFind: false の場合も含む）
	return Result{
		Status: http.StatusOK,
		Data:   result,
	}
}

// 受け取るJSONデータを格納するための構造体
type SignupRequest struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
}

func Signup(req *SignupRequest, uuid string) Result {

	// 必要な情報が入っているか確認
	if req.UserName == "" || req.UserID == "" {
		return Result{
			Message: EmptyInfo,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// 性別の値が正しいか確認
	if req.Gender < 0 || req.Gender > 1 {
		return Result{
			Message: "性別が不正な値です",
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// 誕生日の文字列をtime.Time型に変換
	const layout = "2006-01-02"
	birthday, err := time.Parse(layout, req.Birthday)
	if err != nil {
		return Result{
			Message: InvalidDateTime,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// UUIDの重複を確認
	findUUID, _ := models.GetUserByUUID(uuid)
	if findUUID.IsFind {
		return Result{
			Message: AlreadyUUID,
			Status:  http.StatusConflict,
			Data:    nil,
		}
	}

	// UserIDの重複を確認
	findId, _ := models.GetUserByID(req.UserID)
	if findId.IsFind {
		return Result{
			Message: "ユーザーIDが重複しています",
			Status:  http.StatusConflict,
			Data:    nil,
		}
	}

	// DBにユーザー情報を保存
	result, err := models.Signup(uuid, req.UserName, req.UserID, req.Gender, birthday)
	if err != nil {
		fmt.Println("Database error:", err)
		return Result{
			Message: UnexpectedError,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	//成功時
	return Result{
		Message: "ユーザー情報を保存しました",
		Status:  http.StatusOK,
		Data:    result,
	}
}
