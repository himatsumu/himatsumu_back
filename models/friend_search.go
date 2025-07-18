package models

import (
	"errors"

	"gorm.io/gorm"
)

//フレンドUUIDが存在するか確認する
func FriendSearchByUuid(friendUuid string) (bool) {

	fFriend := Friend{}

	find_result := dbconn.Where(&Friend{FriendUUID: friendUuid}).First(&fFriend)

	if err := find_result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
	