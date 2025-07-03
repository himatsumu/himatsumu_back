package services

//エンドポイントをテストする為のファイル

import (
	"app/models"
	"log"
)

func Debug(user []models.User) {
	//キャラクターを作るサンプル
	chara,err :=models.CreateCharacter()
	log.Println(chara)

	log.Println(user[0].UserUUID, user[1].UserUUID)
	// フレンド申請送信1
	result1 := SendRequest(user[0].UserUUID, user[1].UserUUID)
	if result1.Status != 201 {
		log.Println(result1.Message)
	}
	log.Println("送信結果1", result1.Data)

	// フレンド申請送信2
	result2 := SendRequest(user[3].UserUUID, user[1].UserUUID)
	if result2.Status != 201 {
		log.Println(result2.Message)
	}
	log.Println("送信結果2", result2.Data)

	// 受信済み取得
	result3 := GetRequest(user[1].UserUUID)
	if result3.Status != 200 {
		log.Println(result3.Message)
	}
	maps := result3.Data.([]FriendRequest)
	log.Println("受信結果", maps)

	//拒否
	result4 := ChangeRequestStatus(maps[0].ReqID, maps[0].ReceverId, 2)
	if result4.Status != 200 {
		log.Println(result4.Message)
	}
	log.Println("拒否", result4.Data)

	log.Println(maps[1].ReceverId)	
	log.Println(user[3].UserUUID)
	log.Println(maps[1])

	//承認
	result5 := FriendRecord(maps[1].ReqID,user[3].UserUUID, maps[1].ReceverId)
	if result5.Status != 201 {
		log.Println(result5.Message)
	}
	log.Println("承認", result5.Data)

	//受信済み取得
	result6 := GetFriends(user[1].UserUUID)
	if result6.Status != 200 {
		log.Println(result6.Message)
	}
	log.Println(result6)

	maps2 := result5.Data.(Data)

	log.Println(maps[1].ReceverId)
	log.Println(user[3].UserUUID)
	//クエストクリア
	err = models.QuestCompleted(maps[1].ReceverId,maps2.friendId)
	if err != nil {
		log.Println(err)
	}

	err = models.QuestCompleted(user[3].UserUUID,maps2.friendId)
	if err != nil {
		log.Println(err)
	}

	result7 := IsQuest(maps2.friendId)
	if result7.Status != 200 {
		log.Println(result7.Message)
	}
}
