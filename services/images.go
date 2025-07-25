package services

import (
	// "app/models"
	// "errors"
	// "fmt"
	"mime/multipart"
)

func UploadImage(uid string, file *multipart.FileHeader) error {
    // // レシピ取得
    // user, err := models.GetUserByUuid(uid)
    // if err != nil {
    //     return err
    // }

    // // 空以外の場合
    // if recipe.Image != "" {
    //     return errors.New("Image already exists")
    // }

    // recipe.Image = "https://makeck.tail6cf7b.ts.net:8030/recipe/images/" + uid + ".jpg"

    // // レシピを更新する
    // if err := models.Recipe_Update(recipe); err != nil {
    //     return err
    // }

    // // 画像を保存する
    // filename := fmt.Sprintf("./statics/%s.jpg", uid)
    // if err := saveImage(file, filename); err != nil {
    //     return err
    // }

    return nil
}