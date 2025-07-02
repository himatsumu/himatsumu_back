package services

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// キャラクターの名前を生成
func GenName() Result {
	//JSONファイルを読み込む
	Data, err := ioutil.ReadFile("./configs/name.json")
	if err != nil {
		return Result{
			Message: CouldNotGenerateName,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	// ペットの名前を格納するスライス
	var petNames []string

	// JSONデータをデコードする
	err = json.Unmarshal(Data, &petNames)
	if err != nil {
		return Result{
			Message: CouldNotGenerateName,
			Status:  http.StatusInternalServerError,
			Data:    "",
		}
	}

	//ランダムシードを初期化
	rand.Seed(time.Now().UnixNano())

	// 1から200の間でランダムな数字を生成
	randomNumber := rand.Intn(len(petNames))

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    randomNumber,
	}
}
