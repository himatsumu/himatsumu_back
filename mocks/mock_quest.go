package mocks

import (
	"net/http"
	"time"

	"app/services"

	"github.com/labstack/echo"
)

type createQuests struct {
	schedule string		`json:"schedule"`
	end_time time.Time	`json:"end_time"`
	start_prace string	`json:"start_prace"`
	budget int			`json:"budget"`
	genre int			`json:"genre"`
}

func MockGetQuests(ctx echo.Context) error {
	err :=ctx.Bind(new(createQuests))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": 200,
		"data": echo.Map{
			"stores": []echo.Map{
				{
					"store_name":    "タリーズコーヒー 大阪梅田芝田店",
					"store_address": "大阪府大阪市北区芝田1丁目1-35",
					"start_hours":   "10:00",
					"end_hours":     "20:00",
					"place": echo.Map{
						"lat": 34.705250,
						"lon": 135.497406,
					},
					"types": []string{
						"point_of_interest",
						"cafe",
					},
					"reviews": []string{
						"静かで作業に向いてた",
						"本のセレクトがいい感じ",
						"スイーツもうちょい頑張ってほしい",
						"店員さん優しかった",
        				"電源あって助かった",
					},
				},
				{
					"store_name":    "カフェ バーンホーフ 三番街店",
					"store_address": "大阪府大阪市北区芝田１丁目１−３ 阪急三番街 南館 B2F",
					"start_hours":   "11:00",
					"end_hours":     "20:40",
					"place": echo.Map{
						"lat": 34.704139,
						"lon": 135.499007,
					},
					"types": []string{
						"point_of_interest",
						"cafe",
					},
					"reviews": []string{
						"静かで作業に向いてた",
						"本のセレクトがいい感じ",
						"スイーツもうちょい頑張ってほしい",
						"店員さん優しかった",
        				"電源あって助かった",
					},
				},
				{
					"store_name":    "梅田 阪急三番街 リバーカフェ",
					"store_address": "大阪府大阪市北区芝田１丁目１−３ 阪急三番街南館 地下2階",
					"start_hours":   "11:00",
					"end_hours":     "22:30",
					"place": echo.Map{
						"lat": 34.704139,
						"lon": 135.499007,
					},
					"types": []string{
						"point_of_interest",
						"cafe",
					},
					"reviews": []string{
						"静かで作業に向いてた",
						"本のセレクトがいい感じ",
						"スイーツもうちょい頑張ってほしい",
						"店員さん優しかった",
        				"電源あって助かった",
					},
				},
				{
					"store_name":    "上島珈琲店 阪急三番街店",
					"store_address": "大阪府大阪市北区芝田１丁目１−３ 阪急三番街南館 B1F",
					"start_hours":   "10:00",
					"end_hours":     "21:00",
					"place": echo.Map{
						"lat": 34.704139,
						"lon": 135.499007,
					},
					"types": []string{
						"point_of_interest",
						"cafe",
					},
					"reviews": []string{
						"静かで作業に向いてた",
						"本のセレクトがいい感じ",
						"スイーツもうちょい頑張ってほしい",
						"店員さん優しかった",
        				"電源あって助かった",
					},
				},
			},
			"title": "カフェ",
		},
	})
}

type createQuest struct {
	store_name string	`json:"store_name"`
	store_address string	`json:"store_address"`
	start_hours string	`json:"start_hours"`
	end_hours string	`json:"end_hours"`
	place []string	`json:"place"`
	types []string	`json:"types"`
	reviews []string	`json:"reviews"`
}

func MockCreateQuest(ctx echo.Context) error {
	if err := ctx.Bind(new(createQuest));
	err != nil { 
		return err 
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": 200,
		"data": echo.Map{
			"quest_uuid": "d39f6597-ca73-401c-9a18-4d82ca61d413",
		},
	})
}

type point struct {
	Lat float64
	Lon float64
}

type checkQuest struct {
	quest_uuid string	`json:"quest_uuid"`
	place point		`json:"place"`
}

func MockCheckQuest(ctx echo.Context) error {
	q := new(checkQuest)
	if err := ctx.Bind(q);
	err != nil { 
		return err 
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": 200,
		"data": echo.Map{
			"quest_uuid": q.quest_uuid,
			"point": q.place,
		},
	})
}