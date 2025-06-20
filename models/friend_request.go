package models

import (
	"app/utils"
	"errors"
	"time"
)

//リクエストの存在チェック
func IdRequestfound(Sender_id string, Receiver_id string)(int64,int64) {
	//ネームトークンフィルター
	Sent_filter := FriendReq{}
	Friend_filter := Friend{}

	//既にリクエストが存在していたら1、存在していなかったら0を代入(and)
	request := dbconn.Where(FriendReq{SenderUUID: Sender_id, ReceiverUUID: Receiver_id}).Or(FriendReq{SenderUUID: Receiver_id, ReceiverUUID: Sender_id}).First(&Sent_filter).RowsAffected
	//既にフレンドなら1、フレンドでなかったら0を代入(and)
	friend := dbconn.Where(Friend{UserUUID1: Sender_id, UserUUID2: Receiver_id}).Or(Friend{UserUUID1: Receiver_id, UserUUID2: Sender_id}).First(&Friend_filter).RowsAffected
	
	return request,friend
}

//送信者IDからリクエスト構造体を返す
func ReceivedRequest(Receiver_id string) ([]FriendReq,error){
	//ネームトークンフィルター
	var requests []FriendReq

	//引数と同じ受信側IDの情報を取得
	err := dbconn.Where(FriendReq{ReceiverUUID: Receiver_id}).Find(&requests).Error

	return requests,err
}

// フレンドリクエスト識別子をもとにリクエスト構造体を返す
func IsRequest(uuid string) (FriendReq, error) {

	//ネームトークンフィルター
	request := FriendReq{}

	//UIDが空ならばエラー返す
	if uuid == "" {
		return FriendReq{}, errors.New("UID_does_not_exist")
	}

	//識別子が存在しているか
	result := dbconn.Where(FriendReq{FreReqUUID:uuid}).First(&request)

	//エラーの時
	if result.Error != nil {
		return FriendReq{}, result.Error
	}

	return request, nil
}


//フレンド申請を登録
func SendFriendRequest(Sender_id string,Receiver_id string)(error) {
	//uuid生成
	suid, err := utils.Genid()
	if err != nil {
		return errors.New("uuid generation error")
	}

	//センドトークンの情報
	Stoken := FriendReq{
		FreReqUUID:   suid,        //リクエストの識別子
		SenderUUID:   Sender_id,   //送った側のID
		ReceiverUUID: Receiver_id, //受け取る側のID
		ReqStatus:    0,           ////0:未承認、1:承認、2:拒否、3:取り消し
		ReqUpdateAt:  time.Now(), //最終更新時間
		ReqCreateAt:  time.Now(), //送った時間
	}

	//データベースに書き込む
	if err := dbconn.Create(&Stoken).Error; err != nil {
        return err // エラー処理を追加
    }

	return nil
}

//リクエストの状態を変更
func ChangeRequestStatus(requests FriendReq,reqest int)(error) {
	// フィールドを更新
	requests.ReqStatus = reqest
	requests.ReqUpdateAt = time.Now()

	// 更新されたレコードを保存
	err := dbconn.Updates(&requests).Error

	return err
}