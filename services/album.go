package services

import (
	"app/models"
	"app/utils"
	"net/http"
	"os"
)

//写真フォルダを作る
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

	//フォルダを作る
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

//画像をアップロードする関数
func UplordImg(folderUid string,images string) Result{
	//uuid生成
	uid,err := utils.Genid()
	if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	imgUrl := folderUid + ".png"


    fout, err := os.Create(imgUrl)        
	if err != nil {
        return Result{
        	Message: FolderNotRegistration,
        	Status:  http.StatusInternalServerError,
        	Data:    nil,
        }
    }
	defer fout.Close()

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    uid,
	}
}