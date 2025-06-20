package services

import (
	"app/models"
	"errors"
	"log"
	"net/http"
	"time"
)

const  (
	notaccept = 0   //未承認
	accept = 1      //承認
	reject = 2		//拒否
	cancel = 3		//送信キャンセル
)

// フレンド申請送信
func SendRequest(Sender_id string, Receiver_id string) result {

	//同じユーザーの場合
	if Sender_id == Receiver_id {
		return result{
			status: http.StatusBadRequest,
			err:    errors.New(SameUser),
		}
	}

	//ユーザーが存在するかチェック
	uresult1, err := models.GetUserByID(Sender_id)
	uresult2, err := models.GetUserByID(Receiver_id)

	if !uresult1.IsFind || !uresult2.IsFind {
		return result{
			status: http.StatusNotFound,
			err:    errors.New(UserNotFound),
		}
	}

	//リクエストとフレンドであるか検索する
	request, friend := models.IdRequestfound(Sender_id, Receiver_id)

	//既にリクエストが存在している場合
	if request != 0 {
		return result{
			status: http.StatusBadRequest,
			err:    errors.New(AlreadySent),
		}
	}

	//既にフレンドである場合
	if friend != 0 {
		return result{
			status: http.StatusBadRequest,
			err:    errors.New(AlreadyFriends),
		}
	}

	//データベースに書き込み
	err = models.SendFriendRequest(Sender_id, Receiver_id)

	if err != nil {
		log.Println(err)
		return result{
			status: http.StatusInternalServerError,
			err:    errors.New(FriendRegistrationFailed),
		}
	}

	return result{
		status: http.StatusCreated,
		err:    nil,
	}

}

//フレンドリクエストの戻り値
type FriendRequest struct {
	ReqID      string
	UserID     string
	SenderName string
	ReqTime    time.Time
}

// 受信済み取得
func GetRequest(Receiver_id string) ([]FriendRequest, result) {
	//引数をもとにリクエスト構造体を返す
	requests,err := models.ReceivedRequest(Receiver_id)
	if err != nil {
		return []FriendRequest{},
		result{
			status: http.StatusNotFound,
			err:    errors.New(RequestNotFound),
		}
	}

	//map宣言
	maps := []FriendRequest{}

	for _, request := range requests {
		//IDを元に送信者側の情報取得
		sender, err := models.GetUserByID(request.SenderUUID)

		//ユーザー情報取得に失敗
		if err != nil {
			return []FriendRequest{},
			result{
				status: http.StatusInternalServerError,
				err:    errors.New(UserInfoFailed),
			}
		}

		maps = append(maps, FriendRequest{
			ReqID:      request.FreReqUUID,
			UserID:     request.SenderUUID,
			SenderName: sender.UserData.UserName,
			ReqTime:    request.ReqCreateAt,
		})

	}
	return maps,result{
		status: http.StatusOK,
		err:    nil,
	}
}

//リクエスト状態を変更する
func ChangeRequestStatus(ruid string,ReceiverUUID string,status int)result{
	//リクエストが存在しているか
	request, err := models.IsRequest(ruid)

	//リクエストでエラーならば
	if err != nil || request.ReqStatus != 0 {
		return result{
			status: http.StatusBadRequest,
			err:    errors.New(Incorrectrequesterror),
		}
	}

	//受信者側が一致していない時
	if request.ReceiverUUID != ReceiverUUID {
		return result{
			status: http.StatusBadRequest,
			err:    errors.New(UserMismatchExisting),
		} 
	}

	//フレンドリクエストの状態を変更(statusはint)
	err = models.ChangeRequestStatus(request,status)

	// エラーチェック
	if err != nil {
		log.Println(err)
		return result{
			status: http.StatusInternalServerError,
			err:    errors.New(UserMismatchExisting),
		}
	}

	return result{
		status: http.StatusOK,
		err:    nil,
	}
}
