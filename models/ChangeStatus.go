package models

import "time"

// リクエストの状態を変更
func ChangeRequestStatus(requests FriendReq, Status int) error {
	// フィールドを更新
	requests.ReqStatus = Status
	requests.ReqUpdateAt = time.Now()

	// 更新されたレコードを保存
	err := dbconn.Updates(&requests).Error

	return err
}
