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

	return charaId,nil
}	