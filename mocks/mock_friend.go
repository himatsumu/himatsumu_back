package mocks

import (
	"net/http"

	"github.com/labstack/echo"
)

func MockGetFriends(ctx echo.Context) error {
	
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": 200,
		"data": echo.Map{
			"friends": []echo.Map{
				{
					"user_id": "ayaka",
					"user_name": "あやか",
					"face_image": "/assets/kuma_face.png",
				},
				{
					"user_id": "ruruirui",
					"user_name": "るい",
					"face_image": "/assets/panda_face.png",
				},
				{
					"user_id": "hikaringo",
					"user_name": "ひかり",
					"face_image": "/assets/kuma_face.png",
				},
			},
		},
	})
}

func MockGetFriendById(ctx echo.Context) error {
	userId := ctx.Param("id")
	switch userId {
		case "ayaka":
			return ctx.JSON(http.StatusOK, echo.Map{
				"status": 200,
				"data": echo.Map{
					"user_name": "あやか",
					"point": 100,
					"meet_count": 6,
					"quest_count": 16,
					"exp": 80,
					"level": 5,
					"chara_uuid": "d39f6597-ca73-401c-9a18-4d82ca61d413",
					"chara_name": "ブラウン",
					"chara_image": "/assets/kuma.svg",
					"cos_image": "/assets/kuma_cos.svg",
					"acce_image": "/assets/kuma_acce.svg",
				},
			})
		case "ruruirui":
			return ctx.JSON(http.StatusOK, echo.Map{
				"status": 200,
				"data": echo.Map{
					"user_name": "るい",
					"point": 40,
					"meet_count": 2,
					"quest_count": 6,
					"exp": 80,
					"level": 2,
					"chara_uuid": "d39f6597-ca73-401c-9a18-4d82ca61d413",
					"chara_name": "ブラウン",
					"chara_image": "/assets/kuma.svg",
					"cos_image": "/assets/kuma_cos.svg",
					"acce_image": "/assets/kuma_acce.svg",
				},
			})
		case "hikaringo":
			return ctx.JSON(http.StatusOK, echo.Map{
				"status": 200,
				"data": echo.Map{
					"user_name": "ひかり",
					"point": 1000,
					"meet_count": 10,
					"quest_count": 50,
					"exp": 80,
					"level": 8,
					"chara_uuid": "d39f6597-ca73-401c-9a18-4d82ca61d413",
					"chara_name": "ブラウン",
					"chara_image": "/assets/kuma.svg",
					"cos_image": "/assets/kuma_cos.svg",
					"acce_image": "/assets/kuma_acce.svg",
				},
			})
		default:
			return ctx.JSON(http.StatusOK, echo.Map{
				"status": 200,
				"data": echo.Map{
					"user_name": "あやか",
					"point": 100,
					"meet_count": 6,
					"quest_count": 16,
					"exp": 80,
					"level": 5,
					"chara_uuid": "d39f6597-ca73-401c-9a18-4d82ca61d413",
					"chara_name": "ブラウン",
					"chara_image": "/assets/kuma.svg",
					"cos_image": "/assets/kuma_cos.svg",
					"acce_image": "/assets/kuma_acce.svg",
				},
			})
	}
}