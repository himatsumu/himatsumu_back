package models

import (
	"gorm.io/gorm"
)

type CostumeResult struct {
	CosName string
	CosURL  string
}

// コスチュームUUIDからコスチュームのURLと名前を返す
func GetCostumeByUUID(cosUUID string) (CostumeResult, error) {

	var costume Costume

	result := dbconn.Where(&Costume{CosUUID: cosUUID}).First(&costume)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return CostumeResult{}, gorm.ErrRecordNotFound
		}
		return CostumeResult{}, result.Error
	}

	costumeResult := CostumeResult{
		CosName: costume.CosName,
		CosURL:  costume.CosURL,
	}

	return costumeResult, nil
}
