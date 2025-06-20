package services

import (
	
)

type result struct {
	status int
	err    error
}

const (
	AlreadySent = "すでにリクエストを送信しています"
	SameUser = "同一ユーザーです"
	AlreadyFriends = "既にフレンドです"
	UserNotFound = "ユーザーが見つかりません"
	FriendRegistrationFailed = "フレンドを登録できませんでした"
	RequestNotFound  = "リクエストが存在しませんでした"
	UserInfoFailed = "ユーザー情報取得に失敗しました"
	Incorrectrequesterror = "フレンドリクエストが無効です"
	UserMismatchExisting = "受信者側が一致していません"
	
)