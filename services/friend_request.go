package services

import (
	"app/models"
	"log"
	"net/http"
	"time"
)

// フレンド申請送信
func SendRequest(Sender_id string, Receiver_id string) Result {
	//同じユーザーの場合
	if Sender_id == Receiver_id {
		return Result{
			Message: SameUser,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	//ユーザーが存在するかチェック
	uresult1, err := models.GetUserByID(Sender_id)
	uresult2, err := models.GetUserByID(Receiver_id)

	if !uresult1.IsFind || !uresult2.IsFind {
		return Result{
			Message: UserNotFound,
			Status:  http.StatusNotFound,
			Data:    "",
		}
	}

	//リクエストとフレンドであるか検索する
	request, friend := models.IdRequestfound(Sender_id, Receiver_id)

	//既にリクエストが存在している場合
	if request != 0 {
		return Result{
			Message: AlreadySent,
			Status:  http.StatusBadRequest,
			Data:    "",
		}
	}

	//既にフレンドである場合
	if friend != 0 {
		return Result{
			Message: AlreadyFriends,
			Status:  http.StatusBadRequest,
			Data:    "",
		}
	}
	//データベースに書き込み
	fuid, err := models.SendFriendRequest(Sender_id, Receiver_id)

	if err != nil {
		log.Println(err)
		return Result{
			Message: FriendRegistrationFailed,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusCreated,
		Data:    fuid,
	}
}

// フレンドリクエストの戻り値
type FriendRequest struct {
	ReqID      string
	SenderId   string
	ReceverId  string
	SenderName string
	ReqTime    time.Time
}

// 受信済み取得
func GetRequest(Receiver_id string) Result {
	//引数をもとにリクエスト構造体を返す
	requests, err := models.ReceivedRequest(Receiver_id)
	if err != nil {
		return Result{
			Message: RequestNotFound,
			Status:  http.StatusNotFound,
			Data:    "",
		}
	}

	//map宣言
	maps := []FriendRequest{}

	for _, request := range requests {
		//IDを元に送信者側の情報取得
		sender, err := models.GetUserByID(request.SenderUUID)

		//ユーザー情報取得に失敗
		if err != nil {
			return Result{
				Message: UserInfoFailed,
				Status:  http.StatusInternalServerError,
				Data:    "",
			}
		}

		maps = append(maps, FriendRequest{
			ReqID:      request.FreReqUUID,
			SenderId:   request.SenderUUID,
			ReceverId:  request.ReceiverUUID,
			SenderName: sender.UserData.UserName,
			ReqTime:    request.ReqUpdateAt,
		})
	}

	log.Println("map", maps)

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    maps,
	}
}

// リクエスト状態を変更する
func ChangeRequestStatus(ruid string, ReceiverUUID string, Status int) Result {
	//リクエストが存在しているか
	request, err := models.IsRequest(ruid)

	//リクエストでエラーならば
	if err != nil || request.ReqStatus != 0 {
		return Result{
			Message: Incorrectrequesterror,
			Status:  http.StatusBadRequest,
			Data:    "",
		}
	}

	//受信者側が一致していない時
	if request.ReceiverUUID != ReceiverUUID {
		return Result{
			Message: UserMismatchExisting,
			Status:  http.StatusBadRequest,
			Data:    "",
		}
	}

	//フレンドリクエストの状態を変更(statusはint)
	err = models.ChangeRequestStatus(request, Status)

	// エラーチェック
	if err != nil {
		log.Println(err)
		return Result{
			Message: UserMismatchExisting,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    nil,
	}
}
