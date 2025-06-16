// 様々な要素からユーザーデータを取得する関数があるファイル
package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//検索結果
type FindResult struct{
	IsFind   bool
	UserData User
}

//ユーザーIDを元にユーザーデータを返す
func GetUser_ByID(uid string) (FindResult,error) {

	//空のユーザを作成
	fusers := User{}

	//結果(見つかったらtrue)
	result := FindResult{IsFind: false}

	//空文字のとき
	if uid == "" {
		return result, errors.New("userid is empty")
	}

	//ユーザを取得する
	// find_result := dbconn.Preload(clause.Associations).First(&fuser,&User{UserUUID: uid})
	find_result := dbconn.Where(&User{UserUUID: uid}).Find(&fusers)

	//見つからなかった時
	if err := find_result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return result, gorm.ErrRecordNotFound
	}

	//見つかった時
	result.IsFind = true

	//情報をセットする
	result.UserData = fusers

	return result, nil
}

// ユーザ名でユーザを取得する
func GetUser_ByName(uname string) (FindResult, error) {
	//空のユーザを作成する
	fuser := User{}

	//結果
	result := FindResult{IsFind: false}

	//初期化されていなかったらエラー
	if uname == "" {
		return result, errors.New("username is empty")
	}

	//ユーザを取得する
	find_result := dbconn.Preload(clause.Associations).First(&fuser,&User{UserName: uname})

	//見つからなかった時
	if err := find_result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return result, gorm.ErrRecordNotFound
	}

	//見つかった時
	result.IsFind = true

	//情報をセットする
	result.UserData = fuser

	return result, nil
}

// フレンドリクエスト識別子をもとに、userを2人を返す。存在しなければエラー
func Request(uuid string) (string, string, error) {

	//ネームトークンフィルター
	named_filter := FriendReq{}

	//UIDが空ならばエラー返す
	if uuid == "" {
		return "", "", errors.New("UID_does_not_exist")
	}

	//識別子が存在しているか
	result := dbconn.Where(FriendReq{FreReqUUID:uuid}).First(&named_filter)

	log.Println("named_filter")
	log.Println(result)

	//エラーならば0とエラー型を返す
	if result.Error != nil {
		log.Println("0000")

		return "", "", result.Error
	}

	//Sender_id
	SenderUUID:= named_filter.SenderUUID
	//Receiver_id
	ReceiverUUID := named_filter.ReceiverUUID

	return SenderUUID, ReceiverUUID, nil
}