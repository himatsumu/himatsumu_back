package main

import (
	"app/models"
	"os"
)

func main() {
	// モデル初期化
	models.Init()
	
	// サーバー起動
	mainServer()
}

func DebugModel() {
}

func mainServer() {
	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("GO_URL")))
}
