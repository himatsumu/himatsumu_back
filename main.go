package main

import (
	"app/models"
	"log"
	"os"
)

func main() {

	// モデル初期化
	models.Init()
	
	// サーバー起動
	// mainServer()

	DebugModel()
}

func DebugModel() {

	result := models.Create_Sample_User()
	results := models.Debug(result)
	log.Println(results)
}

func mainServer() {
	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("GO_URL")))
}
