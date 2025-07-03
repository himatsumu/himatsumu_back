package models

import "fmt"
func CreateCharacter() ([]CharaType,error) {
	charaTypes := []CharaType{
		{CharaType: 1, TypeStage: 1, TypeName: "たまご", ImageURL: ""},
		{CharaType: 1, TypeStage: 2, TypeName: "ひび", ImageURL: ""},
		{CharaType: 1, TypeStage: 3, TypeName: "えけちゃん", ImageURL: ""},
		{CharaType: 1, TypeStage: 4, TypeName: "せいじん", ImageURL: ""},
	}

	// データベースに一括で登録
	err := dbconn.Create(&charaTypes).Error
	if err != nil {
		fmt.Println("キャラクターの登録に失敗しました:", err) 
	}
	return charaTypes,err
}