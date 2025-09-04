package main

import (
	"app/models"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {

	//環境変数読み込み
	Init()

	// 公開鍵を保存
	if err := savePublicKeyFromEnv(); err != nil {
		log.Fatal(err)
	}

	// モデル初期化
	err := models.Init()
	log.Println("err", err)

	// サーバー起動
	mainServer()
	// DebugModel()
}

// func DebugModel() {
// 	result := models.CreateSampleUser()
// 	// results := models.Debug(result)
// 	// log.Println("result",results)
// 	services.Debug(result)
// }

func mainServer() {
	// サーバー初期化
	server := InitServer()

	// サーバー起動
	server.Logger.Fatal(server.Start(os.Getenv("GO_URL")))
}

func savePublicKeyFromEnv() error {
	encodedKey := os.Getenv("PUBLIC_KEY_FILE")
	if encodedKey == "" {
		return fmt.Errorf("environment variable PUBLIC_KEY_FILE is not set")
	}

	// Base64文字列をデコードする
	decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return fmt.Errorf("failed to decode base64 public key: %w", err)
	}

	// ディレクトリを作成する
	if err := os.MkdirAll("keys", 0755); err != nil {
		return fmt.Errorf("failed to create keys directory: %w", err)
	}

	// デコードした鍵を書き込む
	filePath := "keys/public.key"
	if err := os.WriteFile(filePath, decodedKey, 0644); err != nil {
		return fmt.Errorf("failed to write public key to file: %w", err)
	}

	log.Printf("Successfully saved public key to %s", filePath)
	return nil
}
