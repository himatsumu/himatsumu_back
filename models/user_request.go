package models

import (
	"time"
)

func Signup (user_uuid string, user_name string, user_id string, gender int, birthday time.Time) (string, error) {
	request := User{
		UserUUID:   user_uuid,
		UserID:     user_id,
		UserName:   user_name,
		Gender:     gender,
		Birthday:   birthday,
		CreateAt:   time.Now(),
	}

	if err := dbconn.Create(&request).Error; err != nil {
		return "", err
	}

	return request.UserID, nil
}