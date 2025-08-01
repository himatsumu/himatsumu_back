package services

import (
	"app/models"
	"app/utils"
	"net/http"
	"os"
)


func CreateFolder(friendUUID string,date string) Result{
	//uuid生成
	uid, err := utils.Genid()
	
	if err != nil  {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	models.CreateFolder(friendUUID,date)
	
	folderPath := os.Getenv("UPLORD_PATH") + uid

    // フォルダを作成（親フォルダも作成）
    err = os.MkdirAll(folderPath, 0755) // 0755はパーミッション
    if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}
	
	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    uid,
	}
}

func UplordImg() {
	
}