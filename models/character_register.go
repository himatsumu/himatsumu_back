package models

import (
	"app/utils"
	"errors"
	"time"
	"math/rand"
)

//キャラクター作成
func RegisterCharacter(frienfId string)(string,error) {
	//uuid生成
	fuid, err := utils.Genid()
	if err != nil {
		return "",errors.New("uuid generation error")
	}

	//キャラクター生成(キャラクターの数)
	randomInt := rand.Intn(4)

	//トークン生成
	Rtoken := Character{
		CharaUUID:  fuid,
		CharaName:  "",
		CharaType:  randomInt,
		TypeStage:  0,
		Exp:        0,
		Birthday:   time.Now().Format("2006-01-02"), 
		Point:      0,
		CharaImage: "",
		OwnChars:   []OwnCharacter{},
	}

	//データベースに書き込む
	if err := dbconn.Create(&Rtoken).Error; err != nil {
        return "",err // エラー処理を追加
    }

	return "",nil
}

//キャラクターとフレンドの中間テーブル生成
func RegisterOwnCharacter(friendId string,characterId string) error{

	//中間テーブル生成トークン生成
	Ctoken := OwnCharacter{
		FriendUUID: friendId,
		CharaUUID: characterId,
	}

	//データベースに書き込む
	if err := dbconn.Create(&Ctoken).Error; err != nil {
        return err // エラー処理を追加
    }

	return nil 
}