package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// テストデータを投入
	err = Seed(dbconn)
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}

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
		&QuestCheck{},
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

// Seed はデータベースに一連のテストデータを投入する
func Seed(db *gorm.DB) error {
	fmt.Println("Seeding database...")

	// --- テストデータ定義 ---
	const (
		userUUID1   = "11111111-1111-1111-1111-111111111111"
		userUUID2   = "22222222-2222-2222-2222-222222222222"
		userUUID3   = "33333333-3333-3333-3333-333333333333"
		friendUUID1 = "f1111111-2222-1111-2222-222222222222"
		charaUUID1  = "caaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
		cosUUID1    = "c05bbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"
		questUUID1  = "qceccccc-cccc-cccc-cccc-cccccccccccc"
		meetUUID1   = "medddddd-dddd-dddd-dddd-dddddddddddd"
		picUUID1    = "p1ceeeee-eeee-eeee-eeee-eeeeeeeeeeee"
	)


	// --- データ本体 ---

	userTestData := []User{
		{UserUUID: userUUID1, UserID: "taro_yamada", UserName: "山田 太郎", Gender: 1, Birthday: time.Date(1995, 4, 1, 0, 0, 0, 0, time.UTC), CreateAt: time.Now()},
		{UserUUID: userUUID2, UserID: "hanako_sato", UserName: "佐藤 花子", Gender: 2, Birthday: time.Date(1998, 8, 15, 0, 0, 0, 0, time.UTC), CreateAt: time.Now()},
		{UserUUID: userUUID3, UserID: "jiro_suzuki", UserName: "鈴木 次郎", Gender: 1, Birthday: time.Date(2000, 1, 20, 0, 0, 0, 0, time.UTC), CreateAt: time.Now()},
	}

	friendTestData := []Friend{
		{FriendUUID: friendUUID1, UserUUID1: userUUID1, UserUUID2: userUUID2, Point: 150, LastMeetAt: time.Now().AddDate(0, 0, -7), CreateAt: time.Now().AddDate(0, -1, 0)},
	}

	friendReqTestData := []FriendReq{
		{FreReqUUID: "freq-1111-3333-1111-333333333333", SenderUUID: userUUID1, ReceiverUUID: userUUID3, ReqStatus: 0, ReqUpdateAt: time.Now(), ReqCreateAt: time.Now()},
	}

	charaTypeTestData := []CharaType{
		{CharaType: 1, TypeStage: 1, TypeName: "くま", ImageURL: "/images/kuma.svg"},
		{CharaType: 2, TypeStage: 1, TypeName: "うさぎ", ImageURL: "/images/usagi.svg"},
	}

	characterTestData := []Character{
		{CharaUUID: charaUUID1, CharaName: "くまさん", CharaType: 1, TypeStage: 1, Exp: 0, Birthday: "04-01", CharaImage: "/images/faces/kuma.svg"},
	}

	ownCharacterTestData := []OwnCharacter{
		{FriendUUID: friendUUID1, CharaUUID: charaUUID1},
	}

	cosTypeTestData := []CosType{
		{CosType: 1, TypeName: "服"},
		{CosType: 2, TypeName: "アクセサリー"},
	}

	costumeTestData := []Costume{
		{CosUUID: cosUUID1, CosName: "探偵帽", CosURL: "/images/costumes/detective_hat.png", Point: 50, CosType: 1, CreateAt: time.Now()},
	}

	ownCostumeTestData := []OwnCostume{
		{FriendUUID: friendUUID1, CosUUID: cosUUID1},
	}

	questHistoryTestData := []QuestHistory{
		{QuestUUID: questUUID1, FriendUUID: friendUUID1, StoreName: "梅田カフェ", StoreAdd: "大阪府大阪市北区", Reviews: []string{"コーヒーが美味しかった", "景色が良い"}, Keywords: []string{"カフェ", "大阪"}, StoType: []string{"cafe", "restaurant"}, Possible: 1, CreateAt: time.Now()},
	}

	questCheckTestData := []QuestCheck{
		{QuestUUID: questUUID1, UserUUID: userUUID1, FriendUUID: friendUUID1, CreateAt: time.Now()},
	}

	
	// マスターデータ
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&charaTypeTestData).Error; err != nil { return err }
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&cosTypeTestData).Error; err != nil { return err }
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&characterTestData).Error; err != nil { return err }
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&costumeTestData).Error; err != nil { return err }

	// ユーザー関連
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&userTestData).Error; err != nil { return err }
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&friendTestData).Error; err != nil { return err }
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&friendReqTestData).Error; err != nil { return err }

	// 保有データ
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&ownCharacterTestData).Error; err != nil { return err }
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&ownCostumeTestData).Error; err != nil { return err }

	// 履歴データ
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&questHistoryTestData).Error; err != nil { return err }

	// 履歴に紐づくデータ
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&questCheckTestData).Error; err != nil { return err }

	fmt.Println("Seeding completed successfully.")
	return nil
}