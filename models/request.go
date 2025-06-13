package models

import (
	"app/utils"
	"errors"
	"log"
	"time"
)

// フレンド申請送信
func SendRequest(Sender_id string, Receiver_id string) (string, error) {
	
	//ネームトークンフィルター
	Sent_filter := FriendReq{}
	Friend_filter := Friend{}

	//既にリクエストが存在していたら1、存在していなかったら0を代入(and)
	request := dbconn.Where(FriendReq{SenderUUID: Sender_id, ReceiverUUID: Receiver_id}).Or(FriendReq{SenderUUID: Receiver_id, ReceiverUUID: Sender_id}).First(&Sent_filter).RowsAffected
	//既にフレンドなら1、フレンドでなかったら0を代入(and)
	friend := dbconn.Where(Friend{UserUUID1: Sender_id, UserUUID2: Receiver_id}).Or(Friend{UserUUID1: Receiver_id, UserUUID2: Sender_id}).First(&Friend_filter).RowsAffected

	log.Println(request)
	log.Println(friend)
	//既にリクエストが存在している場合
	if request != 0 {
		return "", errors.New("request_is_already_existing")
	}

	//既にフレンドである場合
	if friend != 0 {
		return "", errors.New("already_friends")
	}

	//uuid生成
	suid, err := utils.Genid()
	if err != nil {
		return "", errors.New("uuid generation error")
	}

	//センドトークンの情報
	Stoken := FriendReq{
		FreReqUUID:   suid,        //リクエストの識別子
		SenderUUID:   Sender_id,   //送った側のID
		ReceiverUUID: Receiver_id, //受け取る側のID
		ReqStatus:    0,           ////0:未承認、1:承認、2:拒否、3:取り消し
		ReqUpdateAt:  time.Time{}, //最終更新時間
		ReqCreateAt:  time.Time{}, //送った時間
	}

	//データベースに書き込む
	if err := dbconn.Create(&Stoken).Error; err != nil {
        return "", err // エラー処理を追加
    }
	
	log.Println("書き込み完了")
	return suid, err

}

// 受信済み取得
func Get_Request(Receiver_id string) (map[string]map[string]string, error) {

	//ネームトークンフィルター
	var named_filter []FriendReq

	//引数と同じ受信側IDを取得
	length := dbconn.Where(FriendReq{ReceiverUUID: Receiver_id}).Find(&named_filter).RowsAffected

	//map宣言
	maps := map[string]map[string]string{}

	//受信した配列の個数がが0の時
	if length == 0 {
		return map[string]map[string]string{}, nil
	}

	//受信した配列をすべてmapに代入
	for i := 0; i < int(length); i++ {
		uinfo, err := GetUser_ByID(named_filter[i].ReceiverUUID)

		//ユーザー情報取得に失敗
		if err != nil {
			continue
		}

		sinfo, err := GetUser_ByID(named_filter[i].SenderUUID)

		//ユーザー情報取得に失敗
		if err != nil {
			continue
		}

		maps[named_filter[i].FreReqUUID] = map[string]string{
			"uname": uinfo.UserData.UserName,                                   //検索した側の名前
			"sname": sinfo.UserData.UserName,                                   //相手側の名前
			"time" : named_filter[i].ReqUpdateAt.Format("2006-01-02 15:04:05"), //リクエストした時間
		}
	}

	return maps, nil
}
