package models

//エンドポイントをテストする為のファイル

import (
	"errors"
	"log"
)

func Debug(user []User) error {
	//ユーザーIDを元にユーザーデータを返す1
	result, err := GetUser_ByID(user[0].UserUUID)

	if err != nil {
		return err
	}
	if !(result.IsFind) {
		return errors.New("ユーザーいませんでした！(id)")
	}
	

	// ユーザ名でユーザを取得する
	result, err = GetUser_ByName(user[0].UserName)
	if !(result.IsFind) {
		return errors.New("ユーザーいませんでした！(名前)")
	}
	if err != nil {
		return err
	}
	log.Println(result.UserData)

	// フレンド申請送信1
	suid1, err := SendRequest(user[0].UserUUID, user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Println("申請1",suid1)

	// フレンド申請送信2
	suid2, err := SendRequest(user[3].UserUUID, user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Println("申請2",suid2)

	// 受信済み取得
	results, err := Get_Request(user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Println("受信",results)


	aa,err := Rejection(results[0]["id"],user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Print(aa)


	return errors.New("全部OK")
}

