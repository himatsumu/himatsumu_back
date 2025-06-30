package services

import (
	"app/models"
	"net/http"
)

type returnUser struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
}
// 名前を元にユーザー検索
func GetUsersByName(name string) Result {

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
			Id:   results.UserData.UserID,
			Name: results.UserData.UserName,
		},
	}
}

// 名前を元にユーザー検索
func GetUsersById(id string) Result {

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
			Id:   results.UserData.UserID,
			Name: results.UserData.UserName,
		},
	}
}
