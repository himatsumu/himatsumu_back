package services

//エンドポイントをテストする為のファイル

import (
	"app/models"
	"log"
)


func Debug(user []models.User)  {
	log.Println(user[0].UserUUID,user[1].UserUUID)
	// フレンド申請送信1
	result1 := SendRequest(user[0].UserUUID, user[1].UserUUID)
	if result1.status != 200 {
		log.Println(result1.message)
	}
	log.Println("送信結果1",result1.data)
	
	// フレンド申請送信2
	result2 := SendRequest(user[3].UserUUID, user[1].UserUUID)
	if result2.status != 200 {
		log.Println(result2.message)	
	}
	log.Println("送信結果2",result2.data)	
	
	// 受信済み取得
	result3 := GetRequest(user[1].UserUUID)
	if result3.status == 100 {
		log.Println(result3.message)
	}
	maps := result3.data.([]FriendRequest);
	log.Println("受信結果",maps)  
	
	//拒否
	result4 := ChangeRequestStatus(maps[0].ReqID,maps[0].ReceverId,2)
	if result4.status ==  100{
		log.Println(result4.message)
	}
	log.Println("拒否",result4.data)

	//承認
	result5 := ChangeRequestStatus(maps[1].ReqID,maps[1].ReceverId,1)
	if result5.status ==  200{
		log.Println(result5.message)
	}
	log.Println("承認",result5.data)

	//受信済み取得
	result6 := GetFriends(user[1].UserUUID)
	if result6.status == 100 {
		log.Println(result6.message)
	}
	log.Println(result6)

	result7 := GenName()
	_=result7
}

