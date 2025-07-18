package models

import (
	"app/utils"
	"time"
)

type CreateQuestRequest struct {
	FriendUUID string `json:"friend_uuid"`
	StoreName string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	StoreType []string `json:"types"`
	Reviews	[]string `json:"reviews"`
	StorePlace Point `json:"store_place"`
}

// 地点を表す構造体
type Point struct {
	Lat float64 `json:"lat"`	// 緯度（度数法）
	Lon float64 `json:"lon"`	// 経度（度数法）
}

func CreateQuest(userUuid string, req CreateQuestRequest) (string, error) {

	questUUID, err := utils.Genid()
	if err != nil {
		return "", err
	}
	
	quest := QuestHistory{
		QuestUUID:  questUUID,
		FriendUUID: req.FriendUUID,
		StoreName:  req.StoreName,
		StoreAdd: req.StoreAddress,
		StoType: req.StoreType,
		Reviews: req.Reviews,
		StorePlace: req.StorePlace,
		Possible : 0,
		CreateAt:   time.Now(),
	}

	if err := dbconn.Create(&quest).Error; err != nil {
		return "", err // エラー処理を追加
	}

	return questUUID, nil
}

// クエスト達成処理
func QuestCompleted(UserId string, FriendId string) error {
	quest := QuestCheck{
		UserUUID:   UserId,
		FriendUUID: FriendId,
		CreateAt:   time.Now(),
	}

	//データベースに書き込む
	if err := dbconn.Create(&quest).Error; err != nil {
		return err // エラー処理を追加
	}

	return nil
}

// 完了済みクエストの件数を取得
func QuestCount(frienduuid string) (int64, error) {

	//フレンドIDに紐づいている＆現在時間より3分引いた時間内にDBに登録されている件数
	var count int64

	//現在の時間
	currentTime := time.Now()

	//3分前の時間を計算
	threeMinutesAgo := currentTime.Add(-3 * time.Minute)

	err := dbconn.Model(&QuestCheck{FriendUUID: frienduuid}).Where("create_at BETWEEN ? AND ?", threeMinutesAgo, currentTime).Count(&count).Error;

	if err != nil {
		return 0,err
	}
	return count ,nil
}


//完了済みのQuestHistoryテーブルに登録(ここupdateにする)
func QuestsRecorded(frienduuid string)(string,error){

	//uuid生成
	uuid,err := utils.Genid()
	if err != nil {
		return "",err
	}

	history := QuestHistory{
		QuestUUID:  uuid,
		FriendUUID: frienduuid,
		StoreName:  "",
		StoreAdd:  "",
		StoType:    []string{},
		Possible:   0,
		CreateAt:   time.Time{},
	}

	//データベースに書き込む
	if err := dbconn.Create(&history).Error; err != nil {
        return "",err // エラー処理を追加
    }
	
	return uuid,nil
}	