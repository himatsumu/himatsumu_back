package services

import (
	"app/models"
	"app/utils"
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// 写真フォルダを作る
func CreateFolder(friendUUID string, date string) Result {
	//uuid生成
	uid, err := utils.Genid()
	if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	folderPath := os.Getenv("ALBUM_PATH") + friendUUID

	// 親フォルダがなければ作成）
	err = os.MkdirAll(folderPath, 0755) // 0755はパーミッション
	if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	//フォルダを作る
	err = models.CreateFolder(friendUUID, date)
	if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	folderPath = folderPath + "/" + date

	// 日付フォルダを作成
	err = os.MkdirAll(folderPath, 0755) // 0755はパーミッション
	if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    uid,
	}
}

// 画像をアップロードする関数
func UplordImg(friendUUID string, date string, file *multipart.FileHeader) Result {
	//uuid生成
	uid, err := utils.Genid()
	if err != nil {
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}
	folderPath := os.Getenv("ALBUM_PATH") + friendUUID

	imgUrl := folderPath + "/" + date + "/" + uid + ".png"

	// 画像を保存する
	if err := saveImage(file, imgUrl); err != nil {
		log.Print(err)
		return Result{
			Message: FolderNotRegistration,
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	return Result{
		Message: "",
		Status:  http.StatusOK,
		Data:    uid,
	}
}

// 画像を保存する関数
func saveImage(fileHeader *multipart.FileHeader, filename string) error {
	// ファイルを開く
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// 画像をデコード
	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}

	// 画像を保存
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// png形式で保存
	if err := png.Encode(outFile, img); err != nil {
		return err
	}

	return nil
}

type FileInfo struct {
	Name string `json:"name"` // ファイル名
}

type AlbumData struct {
	Date  string     `json:"date"`  // 日付（YYYY-MM-DD形式）
	Album []FileInfo `json:"album"` // その日のファイル一覧
}

// アルバム取得
func GetAlbums(uuid string) Result {
	folderPath := os.Getenv("ALBUM_PATH") + uuid
	// 日付をキーとするマップでファイルをグループ化
	albumMap := make(map[string][]FileInfo)

	// まず、日付フォルダを直接読み込む方法を試す
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		log.Printf("Error reading root folder: %v", err)
		return Result{
			Message: "Error reading root folder",
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	// 各日付フォルダを処理
	for _, entry := range entries {
		entryName := entry.Name()

		if entry.IsDir() {
			// 日付形式かどうかの確認（YYYY-MM-DD形式）
			isDateFormat := len(entryName) == 10 && strings.Count(entryName, "-") == 2

			if isDateFormat {
				dateFolderPath := filepath.Join(folderPath, entryName)

				// 日付フォルダ内のファイルを読み込み
				dateFiles, err := os.ReadDir(dateFolderPath)
				if err != nil {
					// エラーの場合は空のファイルリストで追加
					albumMap[entryName] = []FileInfo{}
					continue
				}

				// ファイルのみを抽出
				var files []FileInfo
				for _, file := range dateFiles {
					if !file.IsDir() {
						files = append(files, FileInfo{
							Name: file.Name(),
						})
					}
				}

				// ファイルがなくても日付フォルダとして追加
				albumMap[entryName] = files
			}
		}
	}

	// マップをスライスに変換し、日付順でソート
	var albums []AlbumData
	for date, files := range albumMap {
		albums = append(albums, AlbumData{
			Date:  date,
			Album: files,
		})
	}

	// 日付順でソート（新しい順）
	sort.Slice(albums, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02", albums[i].Date)
		dateJ, _ := time.Parse("2006-01-02", albums[j].Date)
		return dateI.After(dateJ)
	})

	return Result{
		Message: "Success",
		Status:  http.StatusOK,
		Data:    albums,
	}
}
