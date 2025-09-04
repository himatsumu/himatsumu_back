package models

import (
	"app/utils"
	"errors"
	"time"
)

//フレンド登録
func FriendRecord(Sender_id string,Receiver_id string)(string,error){
	//uuid生成
	fuid, err := utils.Genid()
	if err != nil {
		return "",errors.New("uuid generation error")
	}

	//フレンドトークンの情報
	Stoken := Friend{
		FriendUUID:  fuid,
		UserUUID1:   Sender_id,
		UserUUID2:   Receiver_id,
		CreateAt:    time.Now(),
		OwnChars:    []OwnCharacter{},
		OwnCostumes: []OwnCostume{},
		QuestHis:    []QuestHistory{},
	}

	//データベースに書き込む
	if err := dbconn.Create(&Stoken).Error; err != nil {
        return "",err // エラー処理を追加
    }
	
	return fuid,nil
}

