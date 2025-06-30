package services

import (
	"app/models"
	"net/http"
)

// 戻り値のデータ
type Data struct {
	friendId string
	charaId  string
}

// フレンドになる時の一連の処理
func FriendRecord(ruid string, Sender_id string, Receiver_id string) Result {
	//リクエストの存在チェック
	request, err := models.IsRequest(ruid)
	if err != nil {
		return Result{
			Message: RequestNotFound,
			Status:  http.StatusNotFound,
			Data:    "",
		}
	}

	//与えられたユーザーが違う場合
	if request.ReceiverUUID != Receiver_id || request.SenderUUID != Sender_id {
		return Result{
			Message: UserMismatchExisting,
			Status:  http.StatusBadRequest,
			Data:    "",
		}
	}

	//フレンドテーブルに登録
	friendId, err := models.FriendRecord(Receiver_id, Sender_id)
	if err != nil {
		return Result{
			Message: FriendRegistrationFailed,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	//フレンドリクエストテーブルの状態変更
	err = models.ChangeRequestStatus(request, accept)
	if err != nil {
		return Result{
			Message: FriendRegistrationFailed,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	//キャラクター登録
	charaId, err := RegisterCharacter(friendId)
	if err != nil {
		return Result{
			Message: CharacterNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusCreated,
		Data: Data{
			friendId: friendId,
			charaId:  charaId,
		},
	}
}

// フレンド一覧取得
func GetFriends(uid string) Result {
	getFriends, err := models.GetUserByID(uid)
	if err != nil {
		return Result{
			Message: UserNotFound,
			Status:  http.StatusNotFound,
			Data:    nil,
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    getFriends,
	}
}
