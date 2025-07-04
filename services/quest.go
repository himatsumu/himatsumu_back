package services

import (
	"app/models"
	"net/http"

	"github.com/tidwall/geodesic"
)

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
 