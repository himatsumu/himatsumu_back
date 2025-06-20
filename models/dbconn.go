package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var dbconn *gorm.DB = nil

// InitDB データベース接続を初期化し、オートマイグレーションを実行
func Init() error {
	// データベースに接続
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	dbconn = db

	// オートマイグレーションを実行
	err = autoMigrate()
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %v", err)
	}

	log.Println("Database connection established and auto-migration completed successfully")
	return nil
}

// autoMigrate すべてのモデルを自動マイグレーション
func autoMigrate() error {
	// マイグレーションするモデルのリスト
	models := []interface{}{
		&User{},
		&Friend{},
		&FriendReq{},
		&CharaType{},
		&Character{},
		&OwnCharacter{},
		&CosType{},
		&Costume{},
		&OwnCostume{},
		&QuestHistory{},
		&StoreType{},
		&MeetHistory{},
		&Picture{},
	}

	// 既存のテーブルを削除
	ReseTable(models)

	// トランザクションでマイグレーションを実行
	err := dbconn.Transaction(func(tx *gorm.DB) error {
		for _, model := range models {
			if err := tx.AutoMigrate(model); err != nil {
				return err // エラーが発生したらロールバック
			}
		}
		return nil
	})


	if err != nil {
		return fmt.Errorf("auto migration failed: %v", err)
	}

	return nil
}

//テーブルを全て削除する
func ReseTable(models []interface{}) {

    err := dbconn.Migrator().DropTable(models...)
    if err != nil {
        panic("failed to drop table")
    }
}
