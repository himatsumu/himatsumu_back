# 開発環境
## goのinstall
コンテナの中に入ったらgoをinstall
ctrl + shift + p で　go installで検索
全部install

## gormとechoのinstall
go mod init app
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get github.com/labstack/echo
