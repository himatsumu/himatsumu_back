package services

import (
	"app/models"
)

//キャラクター生成
func RegisterCharacter(friendId string)(string,error){
	//キャラクターテーブルを生成
	charaId,err := models.RegisterCharacter(friendId)
    if err != nil {
		return "",err
	}
	
	//キャラクターとフレンドの中間テーブル生成
	err = models.RegisterOwnCharacter(friendId,charaId)
	if err != nil {
		return "",err
	}
	return charaId,nil
}	