package services

import (
	"app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io"
	"bytes"
	"github.com/tidwall/geodesic"
)

type GenerateQuestsRequest struct {
	Schedule string		`json:"schedule"`
	End_time string	`json:"end_time"`
	Start_prace string	`json:"start_prace"`
	Budget int			`json:"budget"`
	Genre string			`json:"genre"`
}

func GenerateQuests(userUuid string, req GenerateQuestsRequest) Result {

	//UUIDを確認
	if userUuid == "" {
		return Result{
			Message: EmptyUUID,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	fmt.Println(req)

	// 必要な情報が入っているか確認
	if req.Schedule == "" || req.End_time == "" || req.Start_prace == "" || req.Genre == "" {
		return Result{
			Message: EmptyInfo,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	_, err := time.Parse("15:04", req.End_time)
	if err != nil {
		return Result{
			Message: InvalidDateTime,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	//Python-serviceコンテナに送るJSONを作成
	jsonData, err := json.Marshal(req)
	if err != nil {
		return Result{
			Message: "JSON変換に失敗しました",
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	reqBody := bytes.NewBuffer(jsonData)
	
	//クエストを生成する
	aiResponse, err := http.Post("http://python-service:8000/auth/quest/recommend" , "application/json", reqBody)

	if err != nil {
		fmt.Println(err)
		return Result{
			Message: NotRecommendQuest,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	defer aiResponse.Body.Close()

	if aiResponse.StatusCode != http.StatusOK {
		return Result{
			Message: NotRecommendQuest,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	body, err := io.ReadAll(aiResponse.Body)
    if err != nil {
        return Result{
            Message: "レスポンスの読み取りに失敗しました",
            Status:  http.StatusInternalServerError,
        }
    }

	var aiResult interface{}
	err = json.Unmarshal(body, &aiResult)
	if err != nil {
		return Result{
			Message: "JSON変描に失敗しました",
			Status:  http.StatusInternalServerError,
		}
	}

	fmt.Println(aiResult)

	return Result{
		Message: "",
		Status:  http.StatusCreated,
		Data:    aiResult,
	}

}

// 地点を表す構造体
type Point struct {
	Lat float64 // 緯度（度数法）
	Lon float64 // 経度（度数法）
}

// クエスト達成の有無
func QuestRegister(userPoint Point,goalPoint Point,friendId string,userUuid string) Result {	
	
	//2点間の距離を求める
	dist := GetDistance(userPoint,goalPoint)

	if (dist > 20) {
		return  Result{
			Message: QuestNotCompleted,
			Status:  http.StatusUnprocessableEntity,
			Data:    dist,
		}
	}
	//クエストチェックテーブルに完了したことを登録
    err := models.QuestCompleted(friendId,userUuid)
	
	if err != nil {
		return Result{
			Message: "",
			Status:  0,
			Data:    nil,
		}
	}

	return Result{
		Message: "",
		Status:  0,
		Data:    nil,
	}
}

// 2地点間の距離を求める
func GetDistance(point1 Point, point2 Point) int64 {
	//離れている距離
	var dist float64

	//緯度経度を元に距離を比較
	geodesic.WGS84.Inverse(point1.Lat, point1.Lon, point2.Lat, point2.Lon, &dist, nil, nil)

	return int64(dist)
}

// クエストが2人とも達成しているか()
func IsQuest(frienduuid string) Result {

	QuestCount,err := models.QuestCount(frienduuid)

	if err != nil {
		return Result{
			Message: QuestNotCompleted,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	if QuestCount != 2 {
		return Result{
			Message: QuestNotCompleted,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	//ここにmodelsでhistoryの更新を呼びだす

	
	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    nil,
	}
}
 