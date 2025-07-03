package middleware

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

// JWTClaims はJWTに含まれるカスタムクレームです
type JWTClaims struct {
	jwt.RegisteredClaims
}

// JWTを検証する
func JWTAuthMiddleware(publicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status": http.StatusUnauthorized, 
					"message": "認証情報がありません",
				})
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims := &JWTClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return publicKey, nil
			})

			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status": http.StatusUnauthorized,
					"message": "認証情報が不正です",
				})
			}
			if !token.Valid {
				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status": http.StatusUnauthorized,
					"message": "トークンの有効期限が切れています",
				})
			}

			// クレームをEchoのコンテキストに保存
			ctx.Set("user_uuid", claims.Subject)

			// 次の処理を呼び出す
			return next(ctx)
		}
	}
}