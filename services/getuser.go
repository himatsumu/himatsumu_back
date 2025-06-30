package services

import (
	"app/models"
	"net/http"
)

// 名前を元にユーザー検索
func GetUsers(name string) Result {

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
		Data:    results,
	}
}
