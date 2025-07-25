package controllers

import(
	"app/services"
	"net/http"

	"github.com/labstack/echo"
)

func UploadImage(ctx echo.Context) error {
    // body 取得
    uid := ctx.Request().Header.Get("uid")

    // 画像取得
    file, err := ctx.FormFile("image")
    if err != nil {
        // error handling
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    // 検索する
    err = services.UploadImage(uid,file)

    // エラー処理
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    return ctx.JSON(http.StatusOK, echo.Map{
        "result": "success",
    })
}