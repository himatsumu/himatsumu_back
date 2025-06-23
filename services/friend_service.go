package services

import (
	"app/models"
	"net/http"
)

//戻り値のデータ
type data struct {
	friendId string
	charaId  string
}

//フレンドになる時の一連の処理
func FriendRecord(ruid string,Sender_id string,Receiver_id string)(result){
	//リクエストの存在チェック
	request,err := models.IsRequest(ruid)
	if err != nil {
		return result{
			message: RequestNotFound,
			status: http.StatusNotFound,
			data:    "",
		}
	}
	
	//与えられたユーザーが違う場合
	if request.ReceiverUUID != Receiver_id || request.SenderUUID != Sender_id {
		return result{
			message: UserMismatchExisting,
			status: http.StatusBadRequest,
			data:    "",
		}
	}

	//フレンドテーブルに登録
	friendId,err := models.FriendRecord(Receiver_id,Sender_id)
	if err != nil {
		return result{	
			message: FriendRegistrationFailed,
			status: http.StatusInternalServerError,
			data:    "",
		}
	}

	//フレンドリクエストテーブルの状態変更
	err = models.ChangeRequestStatus(request,accept)
	if err != nil {
		return result{
			message: FriendRegistrationFailed,
			status: http.StatusInternalServerError,
			data:    "",
		}
	}

	//キャラクター登録
	charaId,err := RegisterCharacter(friendId)
	if err != nil {
		return result{
			message: CharacterNotRegistration,
			status:  http.StatusInternalServerError,
			data:    "",
		}
	}

	return result{
		message: "",
		status:  http.StatusCreated,
		data: data{
			friendId: friendId,
			charaId:  charaId,
		},
	}
}

//フレンド一覧取得
func GetUser(uid string) result{
	getFriends,err := models.GetUserByID(uid)
	if err != nil {
		return result{
			message: UserNotFound,
			status:  http.StatusNotFound,
			data:    nil,
		}
	}

	return result{
		message: "",
		status:  http.StatusOK,
		data:    getFriends,
	}
}
