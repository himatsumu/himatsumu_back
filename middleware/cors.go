package middleware

import (
	"net/http"
	"os"
	"fmt"

	"github.com/labstack/echo"
)

func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			reactPort := os.Getenv("REACT_PORT")
			reactUrl := os.Getenv("REACT_URL")
			// 許可するオリジン
			allowedOrigins := []string{
				fmt.Sprintf("http://0.0.0.0:%s", reactPort),
            	fmt.Sprintf("http://localhost:%s", reactPort),
				reactUrl,
			}
			req := c.Request()
			res := c.Response()
			origin := req.Header.Get(echo.HeaderOrigin)

			// リクエスト元のオリジンが許可リストに含まれているかチェック
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					res.Header().Set(echo.HeaderAccessControlAllowOrigin, origin)
					break
				}
			}

			res.Header().Set(echo.HeaderAccessControlAllowMethods, "GET, POST, PUT, DELETE, OPTIONS")
			res.Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type, Authorization") // JWTのためにAuthorizationヘッダーを許可
			res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")

			// Preflightリクエスト(OPTIONS)の場合は、ここで処理を終了
			if req.Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}
			
			return next(c)
		}
	}
}
