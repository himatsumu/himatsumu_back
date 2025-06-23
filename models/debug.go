package models

//エンドポイントをテストする為のファイル

import (
	"errors"
	"log"
)

func Debug(user []User) error {
	//ユーザーIDを元にユーザーデータを返す1
	result, err := GetUserByID(user[0].UserUUID)

	if err != nil {
		return err
	}
	if !(result.IsFind) {
		return errors.New("ユーザーいませんでした！(id)")
	}
	
	// ユーザ名でユーザを取得する
	result, err = GetUserByName(user[0].UserName)
	if !(result.IsFind) {
		return errors.New("ユーザーいませんでした！(名前)")
	}
	if err != nil {
		return err
	}

	// フレンド申請送信1
	suid1,err := SendFriendRequest(user[0].UserUUID, user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Println(suid1)

	// フレンド申請送信2
	suid2,err := SendFriendRequest(user[3].UserUUID, user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Println(suid2)
	
	// 受信済み取得
	results, err := ReceivedRequest(user[1].UserUUID)
	if err != nil {
		return err
	}
	log.Println("受信",results)

	//拒否
	err = ChangeRequestStatus(results[0],2)
	if err != nil {
		return err
	}
	
	//承認
	err = ChangeRequestStatus(results[1],1)
	if err != nil {
		return err
	}

	// 受信済み取得
	results, err = ReceivedRequest(user[1].UserUUID)
	if err != nil {
		return err
	}
	for _, result := range results {
		log.Println("ステータス",result.ReqStatus)
	}

	return errors.New("全部OK")
}

