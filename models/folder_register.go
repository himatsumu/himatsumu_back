package models

import (
	"app/utils"
	"errors"
	"time"
)


func CreateFolder(FriendUUID string,date string)(error) {
	//uuid生成
	uid, err := utils.Genid()
	if err != nil {
		return errors.New("uuid generation error")
	}

	// 文字列をtime.Timeに変換
    t, err := time.Parse("2006-01-02", date)
    if err != nil {
        return err
    }

	formattedDate := t.Format("2006/01/02")

	//アルバムトークンの情報
	Stoken := Albums{
		AlbumUUID:  uid,
		FriendUUID: FriendUUID,
		AlbumName:  formattedDate,
		AlbumDate:  date,
	}

	//データベースに書き込む
	if err := dbconn.Create(&Stoken).Error; err != nil {
        return err // エラー処理を追加
    }

	return nil
}