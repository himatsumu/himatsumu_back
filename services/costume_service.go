package services

import (
	"app/custom_error"
	"app/models"
	"fmt"
)

func GetCostume(cos_uuid string) (models.CostumeResult, error) {

	// UUIDが空ならばエラーを返す
	if cos_uuid == "" {
		return models.CostumeResult{}, custom_error.NewBadRequestError("コスチュームが選択されていません")
	}

	result, err := models.GetCostumeByUUID(cos_uuid)
	if err != nil {
		fmt.Println("Database error:", err)
		return models.CostumeResult{}, err
	}

	// 結果をそのまま返す
	return result, nil

}
