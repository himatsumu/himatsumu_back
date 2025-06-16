package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//テーブルを全て削除する
func ReseTable(models []interface{}) {

    err := dbconn.Migrator().DropTable(models...)
    if err != nil {
        panic("failed to drop table")
    }
}

//ユーザーのサンプルデータを入れる
func Create_Sample_User() ([]User){
	// サンプルデータの作成
	users := make([]User, 10)
	for i := 0; i < 10; i++ {
		users[i] = User{
			UserUUID:   uuid.New().String(),
			UserID:     fmt.Sprintf("user%d",i+1),
			UserName:   fmt.Sprintf("User%d",i+1),
			Gender:     i % 2, // 偶数は男性、奇数は女性
			Birthday:   time.Date(1990+i, time.January, 1, 0, 0, 0, 0, time.UTC),
			CreateAt:   time.Now(),
		}
	}
	dbconn.Create(&users)

	return users
}
