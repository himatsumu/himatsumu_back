package services

import (
	"app/models"
	"net/http"
)

type returnUser struct {
	Uuid  string      `json:"uuid"`
	Id    string      `json:"id"`
	Name  string      `json:"name"`
}

// 名前を元にユーザー検索
func GetUsersByName(name string) Result {
	
	if name == "" {
		return Result{
			Message: UserNotFound,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}

	results, err := models.GetUserByName(name)

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
		Data:    returnUser{
			Uuid: results.UserData.UserUUID,
			Id:   results.UserData.UserID,
			Name: results.UserData.UserName,
		},
	}
}

// 名前を元にユーザー検索
func GetUsersById(id string) Result {
	
	if id == "" {
		return Result{
			Message: UserNotFound,
			Status:  http.StatusBadRequest,
			Data:    nil,
		}
	}
	results, err := models.GetUserByID(id)

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
		Data:    returnUser{
			Uuid: results.UserData.UserUUID,
			Id:   results.UserData.UserID,
			Name: results.UserData.UserName,
		},
	}
}
