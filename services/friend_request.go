package services

import (
	"app/models"
	"log"
	"net/http"
	"time"
)

// フレンド申請送信
func SendRequest(Sender_id string, Receiver_id string) result {

	//同じユーザーの場合
	if Sender_id == Receiver_id {
		return result{
			message: SameUser,
			status:  http.StatusBadRequest,
			data:    nil,
		}
	}

	//ユーザーが存在するかチェック
	uresult1, err := models.GetUserByID(Sender_id)
	uresult2, err := models.GetUserByID(Receiver_id)

	if !uresult1.IsFind || !uresult2.IsFind {
		return result{
			message: UserNotFound,
			status: http.StatusNotFound,
			data: "",
		}
	}

	//リクエストとフレンドであるか検索する
	request, friend := models.IdRequestfound(Sender_id,Receiver_id)

	//既にリクエストが存在している場合
	if request != 0 {
		return result{
			message: AlreadySent,
			status: http.StatusBadRequest,
			data: "",
		}
	}

	//既にフレンドである場合
	if friend != 0 {
		return result{
			message: AlreadyFriends,
			status: http.StatusBadRequest,
			data: "",
		}
	}
	//データベースに書き込み
	fuid ,err := models.SendFriendRequest(Sender_id, Receiver_id)

	if err != nil {
		log.Println(err)
		return result{
			message: FriendRegistrationFailed,
			status: http.StatusInternalServerError,
			data: "",
		}
	}

	return result{
		message: "",
		status: http.StatusCreated,
		data: fuid,
	}
}

//フレンドリクエストの戻り値
type FriendRequest struct {
	ReqID       string
	SenderId     string
	ReceverId   string
	SenderName string
	ReqTime    time.Time
}

// 受信済み取得
func GetRequest(Receiver_id string) (result) {
	//引数をもとにリクエスト構造体を返す
	requests,err := models.ReceivedRequest(Receiver_id)
	if err != nil {
		return  result{
			message: RequestNotFound,
			status: http.StatusNotFound,
			data: "",
		}
	}

	//map宣言
	maps := []FriendRequest{}

	for _, request := range requests {
		//IDを元に送信者側の情報取得
		sender, err := models.GetUserByID(request.SenderUUID)

		//ユーザー情報取得に失敗
		if err != nil {
			return result{
				message: UserInfoFailed,
				status: http.StatusInternalServerError,
				data: "",
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

	log.Println("map",maps)

	return result{
		message: "",
		status: http.StatusOK,
		data:    maps,
	}
}

//リクエスト状態を変更する
func ChangeRequestStatus(ruid string,ReceiverUUID string,status int)result{
	//リクエストが存在しているか
	request, err := models.IsRequest(ruid)

	//リクエストでエラーならば
	if err != nil || request.ReqStatus != 0 {
		return result{
			message: Incorrectrequesterror,
			status: http.StatusBadRequest,
			data: "",
		}
	}

	//受信者側が一致していない時
	if request.ReceiverUUID != ReceiverUUID {
		return result{
			message: UserMismatchExisting,
			status: http.StatusBadRequest,
			data: "",
		} 
	}

	//フレンドリクエストの状態を変更(statusはint)
	err = models.ChangeRequestStatus(request,status)

	// エラーチェック
	if err != nil {
		log.Println(err)
		return result{
			message: UserMismatchExisting,
			status: http.StatusInternalServerError,
			data: "",
		}
	}

	return result{
		message: "",
		status: http.StatusOK,
		data:    nil,
	}
}

