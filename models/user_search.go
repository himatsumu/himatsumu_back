// 様々な要素からユーザーデータを取得する関数があるファイル
package models

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//検索結果
type FindResult struct{
	IsFind   bool
	UserData *User
}

//ユーザーIDを元にユーザーデータを返す
func GetUserByID(uid string) (FindResult,error) {

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
	result.UserData = &fusers

	return result, nil
}

// ユーザ名でユーザを取得する
func GetUserByName(uname string) (FindResult, error) {
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
	result.UserData = &fuser

	return result, nil
}

func GetUserByUUID(uuid string) (FindResult, error) {
	//空のユーザを作成する
	fuser := User{}
	
	//結果
	result := FindResult{IsFind: false}
	
	//ユーザを取得する
	find_result := dbconn.Preload(clause.Associations).First(&fuser, &User{UserUUID: uuid})
	
	//見つからなかった時 - エラーではなく、IsFind: falseで返す
	if err := find_result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return result, nil // エラーではなくnilを返す
	}
	
	//その他のエラーがあった場合
	if find_result.Error != nil {
		return result, find_result.Error
	}
	
	//見つかった時
	result.IsFind = true
	//情報をセットする
	result.UserData = &fuser
	
	return result, nil
}