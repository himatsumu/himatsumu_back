package models

import (
	"errors"

	"gorm.io/gorm"
)

// フレンドUUIDが存在するか確認する
func FriendSearchByUuid(friendUuid string) bool {

	fFriend := Friend{}

	find_result := dbconn.Where(&Friend{FriendUUID: friendUuid}).First(&fFriend)

	if err := find_result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

type FriendResponse struct {
	FriendUuid     string
	FriendUserUuid string
	FriendName     string
	FaceImg        string
}

func GetFriends(uuid string) ([]FriendResponse, error) {
	var friendRelations []Friend

	// Preloadを使い、関連するUserテーブルのデータを一括で取得する
	err := dbconn.
		Preload("User1").
		Preload("User2").
		// 自分自身がUserUUID1かUserUUID2に含まれるフレンド関係をすべて検索
		Where("\"USER_UUID1\" = ? OR \"USER_UUID2\" = ?", uuid, uuid).
		Find(&friendRelations).Error

	if err != nil {
		return nil, err
	}

	// ---取得したデータをレスポンス形式に変換---
	var response []FriendResponse

	for _, rel := range friendRelations {
		var friendUser User
		// 自分ではない方のユーザー情報を特定する
		if rel.UserUUID1 != uuid {
			friendUser = rel.User1
		} else {
			friendUser = rel.User2
		}

		// レスポンス用のオブジェクトを作成
		respFriend := FriendResponse{
			FriendUuid:     rel.FriendUUID,
			FriendUserUuid: friendUser.UserUUID,
			FriendName:     friendUser.UserName,
			FaceImg:        rel.FaceImg,
		}
		response = append(response, respFriend)
	}

	return response, nil
}

type FriendPoint struct {
	Point       int
}

func GetFriendPoint(uuid string) (Friend, error) {
	var friendPoint Friend
	err := dbconn.Where("\"FRIEND_UUID\" = ?", uuid).First(&friendPoint).Error
	return friendPoint, err
}