package mocks

import (
	"net/http"

	"github.com/labstack/echo"
)

func MockGetQuests(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": 200,
		"data": echo.Map{
			"stores": []echo.Map{
				{
					"store_name":    "タリーズコーヒー 大阪梅田芝田店",
					"store_address": "大阪府大阪市北区芝田1丁目1-35",
					"start_hours":   "10:00",
					"end_hours":     "20:00",
					"lat":           34.705250,
					"lon":           135.497406,
				},
				{
					"store_name":    "カフェ バーンホーフ 三番街店",
					"store_address": "大阪府大阪市北区芝田１丁目１−３ 阪急三番街 南館 B2F",
					"start_hours":   "11:00",
					"end_hours":     "20:40",
					"lat":           34.704139,
					"lon":           135.499007,
				},
				{
					"store_name":    "梅田 阪急三番街 リバーカフェ",
					"store_address": "大阪府大阪市北区芝田１丁目１−３ 阪急三番街南館 地下2階",
					"start_hours":   "11:00",
					"end_hours":     "22:30",
					"lat":           34.704139,
					"lon":           135.499007,
				},
				{
					"store_name":    "上島珈琲店 阪急三番街店",
					"store_address": "大阪府大阪市北区芝田１丁目１−３ 阪急三番街南館 B1F",
					"start_hours":   "10:00",
					"end_hours":     "21:00",
					"lat":           34.704139,
					"lon":           135.499007,
				},
			},
			"title": "カフェ",
		},
	})
}
