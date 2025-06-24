package models

import "time"

//リクエストの状態を変更
func ChangeRequestStatus(requests FriendReq,status int)(error) {
	// フィールドを更新
	requests.ReqStatus = status
	requests.ReqUpdateAt = time.Now()

	// 更新されたレコードを保存
	err := dbconn.Updates(&requests).Error

	return err
}