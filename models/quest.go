package models

import (
	"app/utils"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
	"gorm.io/gorm"
)

type CreateQuestRequest struct {
	FriendUUID   string   `json:"friend_uuid"`
	StoreName    string   `json:"store_name"`
	StoreAddress string   `json:"store_address"`
	StoreType    []string `json:"types"`
	Reviews      []string `json:"reviews"`
	StorePlace   Point    `json:"store_place"`
}

// Point構造体をJSON形式に変換
func (p Point) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// JSON形式をPoint構造体に変換
func (p *Point) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &p)
}

// 地点を表す構造体
type Point struct {
	Lat float64 `json:"lat"` // 緯度（度数法）
	Lon float64 `json:"lon"` // 経度（度数法）
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
		StoreAdd:   req.StoreAddress,
		StoType:    req.StoreType,
		Reviews:    req.Reviews,
		StorePlace: req.StorePlace,
		Possible:   0,
		CreateAt:   time.Now(),
	}

	if err := dbconn.Create(&quest).Error; err != nil {
		return "", err // エラー処理を追加
	}

	return questUUID, nil
}

// クエストの達成場所を取得
func GetQuestPoint(questUUID string) (Point, error) {
	var quest QuestHistory
	if err := dbconn.Where("\"QUEST_UUID\" = ?", questUUID).First(&quest).Error; err != nil {
		return Point{}, err // エラー処理を追加
	}

	return quest.StorePlace, nil
}

// クエスト達成処理
func QuestCompleted(QuestUuid string, UserId string, FriendId string) error {
	quest := QuestCheck{
		QuestUUID:  QuestUuid,
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
func QuestCount(questUuid string) (int64, error) {

	//フレンドIDに紐づいている＆現在時間より3分引いた時間内にDBに登録されている件数
	var count int64

	//現在の時間
	currentTime := time.Now()

	//3分前の時間を計算
	threeMinutesAgo := currentTime.Add(-3 * time.Minute)

	err := dbconn.Model(&QuestCheck{QuestUUID: questUuid}).Where("create_at BETWEEN ? AND ?", threeMinutesAgo, currentTime).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

// 完了済みのQuestHistoryテーブルに登録(ここupdateにする)
func QuestsRecorded(frienduuid string) (string, error) {

	//uuid生成
	uuid, err := utils.Genid()
	if err != nil {
		return "", err
	}

	history := QuestHistory{
		QuestUUID:  uuid,
		FriendUUID: frienduuid,
		StoreName:  "",
		StoreAdd:   "",
		StoType:    []string{},
		Possible:   0,
		CreateAt:   time.Time{},
	}

	//データベースに書き込む
	if err := dbconn.Create(&history).Error; err != nil {
		return "", err // エラー処理を追加
	}

	return uuid, nil
}

func UpdateHistory(questUuid string) error {

	result := dbconn.Model(&QuestHistory{}).Where("\"QUEST_UUID\" = ?", questUuid).Update("\"POSSIBLE\"", 1)

		// エラーが発生した場合
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
}

func UpdateCharaAndPoint(charaUuid string, friendUuid string, exp int, levelUpCount int, point int) error {

	updatesChara := map[string]interface{}{
		"\"EXP\"": exp,
	}

	// 更新対象のカラム名を保持する変数
	var columnToUpdate string

	var chara Character

	if levelUpCount == 1 {
		if err := dbconn.Where("\"CHARA_UUID\" = ?", charaUuid).First(&chara).Error; err != nil {
			return err
		}

		// どの進化日を更新すべきか判定する
		if chara.FirstEvo.IsZero() {
			columnToUpdate = "first_evo"
		} else if chara.SecondEvo.IsZero() {
			columnToUpdate = "second_evo"
		} else if chara.ThirdEvo.IsZero() {
			columnToUpdate = "third_evo"
		} else if chara.FourthEvo.IsZero() {
			columnToUpdate = "fourth_evo"
		} else if chara.FifthEvo.IsZero() {
			columnToUpdate = "fifth_evo"
		} else if chara.SixthEvo.IsZero() {
			columnToUpdate = "sixth_evo"
		}
	}

	
	if columnToUpdate != "" {
		updatesChara[columnToUpdate] = time.Now()
		updatesChara["\"TYPE_STAGE\""] = chara.TypeStage + 1
	}

	result := dbconn.Model(&Character{}).Where("\"CHARA_UUID\" = ?", charaUuid).Updates(updatesChara)
	if result.Error != nil {
		return result.Error
	}

	// フレンドポイントの更新
	result = dbconn.Model(&Friend{}).Where("\"FRIEND_UUID\" = ?", friendUuid).Updates(map[string]interface{}{"\"POINT\"": point})
	if result.Error != nil {
		return result.Error
	}
		
	return nil
}