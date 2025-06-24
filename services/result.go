package services

import (
	
)

type result struct {
    message string  `json:"message"`
    status  int    `json:"status"`
    data    interface{} `"json:data"`
}

const  (
	notaccept = 0   //未承認
	accept = 1      //承認
	reject = 2		//拒否
	cancel = 3		//送信キャンセル
)


const (
	AlreadySent = "すでにリクエストを送信しています"
	SameUser = "同一ユーザーです"
	AlreadyFriends = "既にフレンドです"
	UserNotFound = "ユーザーが見つかりません"
	FriendRegistrationFailed = "フレンドを登録できませんでした"
	RequestNotFound  = "リクエストが存在しませんでした"
	UserInfoFailed = "ユーザー情報取得に失敗しました"
	Incorrectrequesterror = "フレンドリクエストが無効です"
	UserMismatchExisting = "ユーザーが一致していません"
	CharacterNotRegistration = "キャラクターを生成できませんでした"
	CouldNotGenerateName = "名前を生成できませんでした"
	
)