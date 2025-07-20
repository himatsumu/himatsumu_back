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
	"errors"
	"gorm.io/gorm"
	"github.com/jackc/pgx/v5/pgconn"
)

type GenerateQuestsRequest struct {
	Schedule string		`json:"schedule"`
	End_time string	`json:"end_time"`
	Start_prace string	`json:"start_prace"`
	Budget int			`json:"budget"`
	Genre string			`json:"genre"`
}

type QuestUUID struct {
	QuestUUID string `json:"quest_uuid"`
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

	// レスポンスを閉じる
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

	//JSONを構造体に変換
	var aiResult interface{}
	err = json.Unmarshal(body, &aiResult)
	if err != nil {
		return Result{
			Message: "JSON変描に失敗しました",
			Status:  http.StatusInternalServerError,
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusCreated,
		Data:    aiResult,
	}

}

func CreateQuest(userUuid string, req models.CreateQuestRequest) Result {

	//UUIDを確認
	if userUuid == "" {
		return Result{
			Message: EmptyUUID,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// 必要な情報が入っているか確認
	if req.FriendUUID == "" || req.StoreName == "" || req.StoreAddress == "" {
		return Result{
			Message: EmptyInfo,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// フレンドが存在しているか確認
	friendCheck := models.FriendSearchByUuid(req.FriendUUID)

	if !friendCheck {
		return Result{
			Message: NotFriend,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	result, err := models.CreateQuest(userUuid, req)
	if err != nil {
		return Result{
			Message: UnexpectedError,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusCreated,
		Data:    QuestUUID{
			QuestUUID: result,
		},
	}
}

type CheckQuestRequest struct {
	Point models.Point `json:"point"`
	QuestUuid string `json:"quest_uuid"`
	FriendUuid string `json:"friend_uuid"`
}

// クエスト達成の有無
func CheckQuest(userUuid string, req CheckQuestRequest) Result {	
	
	//UUIDを確認
	if userUuid == "" {
		return Result{
			Message: EmptyUUID,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// 必要な情報が入っているか確認
	if req.QuestUuid == "" || req.FriendUuid == "" {
		return Result{
			Message: EmptyInfo,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	// フレンドが存在しているか確認
	friendCheck := models.FriendSearchByUuid(req.FriendUuid)
	if !friendCheck {
		return Result{
			Message: NotFriend,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	//クエストの位置情報を取得する
	goalPoint, err := models.GetQuestPoint(req.QuestUuid)
	if err != nil {
		// レコードが見つからなかった場合
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Result{
				Message: QuestNotFound,
				Status:  http.StatusNotFound,
				Data:    nil,
			}
		}

		// それ以外のエラー
		return Result{
			Message: UnexpectedError,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	//2点間の距離を求める
	dist := GetDistance(req.Point,goalPoint)

	if (dist > 20) {
		return  Result{
			Message: QuestNotCompleted,
			Status:  http.StatusUnprocessableEntity,
			Data:    dist,
		}
	}
	//クエストチェックテーブルに完了したことを登録
    err = models.QuestCompleted(req.QuestUuid, userUuid, req.FriendUuid)
	
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return Result{
				Message: "クエストを達成済みです",
				Status:  http.StatusUnprocessableEntity,
				Data:    nil,
			}
		}

		// それ以外のエラー
		return Result{
			Message: UnexpectedError,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    nil,
	}
}

// 2地点間の距離を求める
func GetDistance(point1 models.Point, point2 models.Point) int64 {
	//離れている距離
	var dist float64

	//緯度経度を元に距離を比較
	geodesic.WGS84.Inverse(point1.Lat, point1.Lon, point2.Lat, point2.Lon, &dist, nil, nil)

	return int64(dist)
}

// クエストが2人とも達成しているか()
func IsQuest(questUuid string) Result {

	QuestCount,err := models.QuestCount(questUuid)

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
 