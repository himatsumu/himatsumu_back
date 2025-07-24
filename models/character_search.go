package models

func GetCharacterDetail(charaUuid string) (Character, error) {
	var chara Character
	if err := dbconn.Where("\"CHARA_UUID\" = ?", charaUuid).First(&chara).Error; err != nil {
		return Character{}, err // エラー処理を追加
	}

	return chara, nil
}
