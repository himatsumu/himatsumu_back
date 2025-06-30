package services

import (
	"app/middleware"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 認証済みユーザーの情報を返します
func (s *UserService) GetAuthenticatedData(c echo.Context) error {
	// ミドルウェアによってコンテキストに保存されたクレーム情報を取得
	claims, ok := c.Get("claims").(*middleware.JWTClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not retrieve user claims from context",
		})
	}

	// レスポンスを作成
	response := map[string]interface{}{
		"Message":      "You have accessed a protected endpoint!",
		"user_uuid":    claims.Subject,
		"token_issuer": claims.Issuer,
	}

	fmt.Println(response)
	// c.JSON() を使ってJSONレスポンスを返す
	return c.JSON(http.StatusOK, response)
}
