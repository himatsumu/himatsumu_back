package services

import (
	"app/custom_error"
	"app/middleware"
	"app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// 認証済みユーザーの情報を返す
func GetAuthenticatedData(c echo.Context) error {
	// ミドルウェアによってコンテキストに保存されたクレーム情報を取得
	claims, ok := c.Get("claims").(*middleware.JWTClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not retrieve user claims from context",
		})
	}

	// レスポンスを作成
	response := map[string]interface{}{
		"message":      "You have accessed a protected endpoint!",
		"user_uuid":    claims.Subject,
		"token_issuer": claims.Issuer,
	}

	fmt.Println(response)
	// c.JSON() を使ってJSONレスポンスを返す
	return c.JSON(http.StatusOK, response)
}

// ユーザー情報を返す
func CheckUser(uuid string) (models.FindResult, error) {
	// ユーザを取得する(なければいないことを返す)
	result, err := models.GetUserByUUID(uuid)
	if err != nil {
		fmt.Println("Database error:", err)
		return models.FindResult{IsFind: false}, err
	}

	// 結果をそのまま返す（IsFind: false の場合も含む）
	return result, nil
}

// 受け取るJSONデータを格納するための構造体
type SignupRequest struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	Gender   int    `json:"gender"`
	Birthday string `json:"birthday"`
}

func Signup(req *SignupRequest, ctx echo.Context) (string, error) {

	// リクエストからユーザー情報を取得
	user_uuid := ctx.Get("user_uuid").(string)

	// 必要な情報が入っているか確認
	if req.UserName == "" || req.UserID == "" {
		// HTTPレスポンスではなく、errorを返す
		return "", custom_error.NewBadRequestError("UserName、またはUserIDが入力されていません")
	}

	// 性別の値が正しいか確認
	if req.Gender < 0 || req.Gender > 1 {
		return "", custom_error.NewBadRequestError("性別が正しく入力されていません")
	}

	// 誕生日の文字列をtime.Time型に変換
	const layout = "2006-01-02"
	birthday, err := time.Parse(layout, req.Birthday)
	if err != nil {
		return "", custom_error.NewBadRequestError("誕生日が正しくない日付形式です")
	}

	// UUIDの重複を確認
	if _, err := models.GetUserByUUID(user_uuid); err != nil {
		return "", custom_error.NewConflictError("UUIDが重複しています")
	}

	// UserIDの重複を確認
	if _, err := models.GetUserByID(req.UserID); err != nil {
		return "", custom_error.NewConflictError("UserIDが重複しています")
	}

	// DBにユーザー情報を保存
	result, err := models.Signup(user_uuid, req.UserName, req.UserID, req.Gender, birthday)
	if err != nil {
		fmt.Println("Database error:", err)
		return "", err
	}

	//成功時
	return result, nil
}
