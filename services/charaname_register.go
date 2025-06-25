package services

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// キャラクターの名前を生成
func GenName() result {
	//JSONファイルを読み込む
	data, err := ioutil.ReadFile("./configs/name.json")
	if err != nil {
		return result{
			message: CouldNotGenerateName,
			status:  http.StatusInternalServerError,
			data:    "",
		}
	}

	// ペットの名前を格納するスライス
	var petNames []string

	// JSONデータをデコードする
	err = json.Unmarshal(data, &petNames)
	if err != nil {
		return result{
			message: CouldNotGenerateName,
			status:  http.StatusInternalServerError,
			data:    "",
		}
	}

	//ランダムシードを初期化
	rand.Seed(time.Now().UnixNano())

	// 1から200の間でランダムな数字を生成
	randomNumber := rand.Intn(len(petNames))

	return result{
		message: "",
		status:  http.StatusOK,
		data:    randomNumber,
	}
}

// キャラクターの名前を設定 
func RegisterCharaname(charauid string,name string) result {
	err := models.RegisterCharacterName(charauid,name)
	if err != nil {
		return result{
			message: CouldNotGenerateName,
			status:  http.StatusInternalServerError,
			data:    nil,
		}
	} 

	return result{
		message: "",
		status:  http.StatusOK,
		data:    nil,
	}
}
