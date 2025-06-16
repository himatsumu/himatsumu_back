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
	UUID string `json:"uuid`
	jwt.RegisteredClaims
}

// JWTを検証する
func JWTAuthMiddleware(publicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is required"})
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
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token: " + err.Error()})
			}
			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is not valid"})
			}

			// クレームをEchoのコンテキストに保存
			c.Set("claims", claims)

			// 次の処理を呼び出す
			return next(c)
		}
	}
}