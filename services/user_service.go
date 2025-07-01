package services

import (
	"app/middleware"
	"app/models"
	"fmt"
	"net/http"

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
