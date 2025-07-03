package models

import (
	"time"
)

// User ユーザーテーブル
type User struct {
	UserUUID   string      `gorm:"primaryKey;column:USER_UUID;type:CHAR(36);not null"`
	UserID     string      `gorm:"column:USER_ID;type:VARCHAR(20);not null;uniqueIndex"`
	UserName   string      `gorm:"column:USER_NAME;type:VARCHAR(20);not null"`
	Gender     int         `gorm:"column:GENDER;type:INT;not null"`
	Birthday   time.Time   `gorm:"column:BIRTHDAY;type:DATE;not null"`
	CreateAt   time.Time   `gorm:"column:CREATE_AT;type:timestamp;not null"`
	Friends    []Friend    `gorm:"foreignKey:UserUUID1;references:UserUUID"`
	FriendReqs []FriendReq  `gorm:"foreignKey:SenderUUID;references:UserUUID"` // 修正
}

// Friend フレンドテーブル
type Friend struct {
	FriendUUID  string         `gorm:"primaryKey;column:FRIEND_UUID;type:CHAR(36);not null"`
	UserUUID1   string         `gorm:"column:USER_UUID1;type:CHAR(36);not null;uniqueIndex:idx_user_uuid_pair"`
	UserUUID2   string         `gorm:"column:USER_UUID2;type:CHAR(36);not null;uniqueIndex:idx_user_uuid_pair"`
	LastMeetAt  time.Time      `gorm:"column:LAST_MEET_AT;type:DATE;"`
	CreateAt    time.Time      `gorm:"column:CREATE_AT;type:timestamp;not null"`
	OwnChars    []OwnCharacter `gorm:"foreignKey:FriendUUID;references:FriendUUID"`
	OwnCostumes []OwnCostume   `gorm:"foreignKey:FriendUUID;references:FriendUUID"`
	QuestHis    []QuestHistory `gorm:"foreignKey:FriendUUID;references:FriendUUID"`
	MeetHis     []MeetHistory  `gorm:"foreignKey:FriendUUID;references:FriendUUID"`
}

// FriendReq フレンドリクエストテーブル
type FriendReq struct {
	FreReqUUID   string    `gorm:"primaryKey;column:FRE_REQ_UUID;type:CHAR(36);not null"`
	SenderUUID   string    `gorm:"column:Sender_UUID;type:CHAR(36);not null"` // 修正
	ReceiverUUID string    `gorm:"column:Receiver_UUID;type:CHAR(36);not null"` // 修正
	ReqStatus    int       `gorm:"column:REQ_STATUS;type:INT;not null"` // 0:未承認、1:承認、2:拒否、3:取り消し
	ReqUpdateAt  time.Time `gorm:"column:REQ_UPDATE_AT;type:timestamp;not null"`
	ReqCreateAt  time.Time `gorm:"column:REQ_CREATE_AT;type:timestamp;not null"`
}

// CharaType キャラクター種別テーブル
type CharaType struct {
	CharaType  int         `gorm:"primaryKey;column:CHARA_TYPE;type:INT;not null"`
	TypeStage  int         `gorm:"primaryKey;column:TYPE_STAGE;type:INT;not null;"`
	TypeName   string      `gorm:"column:TYPE_NAME;type:VARCHAR(20);not null"`
	ImageURL   string      `gorm:"column:IMAGE_URL;type:VARCHAR(50);not null"`
}

// Character キャラクターテーブル
type Character struct {
	CharaUUID  string         `gorm:"primaryKey;column:CHARA_UUID;type:CHAR(36);not null"`
	CharaName  string         `gorm:"column:CHARA_NAME;type:VARCHAR(36);not null"`
	CharaType  int            `gorm:"column:CHARA_TYPE;type:INT;not null"`
	TypeStage  int            `gorm:"column:TYPE_STAGE;type:INT;not null"`
	Exp        int            `gorm:"column:EXP;type:INT;not null;default:0"`
	Birthday   string         `gorm:"column:BARTHDAY;"`
	FirstEvo   time.Time 	  `gorm:"column:FIRST_EVO;type:timestamp;`
	SecondEvo  time.Time 	  `gorm:"column:SECOND_EVO;type:timestamp;`
	ThirdEvo   time.Time 	  `gorm:"column:THIRD_EVO;type:timestamp;`
	FourthEvo  time.Time 	  `gorm:"column:FOURTH_EVO;type:timestamp"`
	FifthEvo   time.Time 	  `gorm:"column:FIFTH_EVO;type:timestamp"`
	SixthEvo   time.Time 	  `gorm:"column:SIXTH_EVO;type:timestamp"`
	Point      int            `gorm:"column:POINT;type:INT;not null;default:0"`
	CharaImage string         `gorm:"column:CHARA_IMAGE;type:VARCHAR(50)"`
	OwnChars   []OwnCharacter `gorm:"foreignKey:CharaUUID;references:CharaUUID"`
}

// OwnCharacter 保有キャラクターテーブル
type OwnCharacter struct {
	FriendUUID string `gorm:"column:FRIEND_UUID;type:CHAR(36);not null;foreignKey:FriendUUID;references:FriendUUID"`
	CharaUUID  string `gorm:"column:CHARA_UUID;type:CHAR(36);not null;foreignKey:CharaUUID;references:CharaUUID"`
}

// CosType 衣装種別テーブル
type CosType struct {
	CosType  int       `gorm:"primaryKey;column:COS_TYPE;type:INT;not null"`
	TypeName string    `gorm:"column:TYPE_NAME;type:VARCHAR(30);not null"`
	Costumes []Costume `gorm:"foreignKey:CosType;references:CosType"`
}

// Costume キャラクター衣装テーブル
type Costume struct {
	CosUUID     string       `gorm:"primaryKey;column:COS_UUID;type:CHAR(36);not null"`
	CosName     string       `gorm:"column:COS_NAME;type:VARCHAR(20);not null"`
	CosURL      string       `gorm:"column:COS_URL;type:VARCHAR(50);not null"`
	Point       int          `gorm:"column:POINT;type:INT;not null;default:0"`
	CosType     int          `gorm:"column:COS_TYPE;type:INT;not null"`
	CreateAt    time.Time    `gorm:"column:CREATE_AT;type:timestamp;not null"`
	OwnCostumes []OwnCostume `gorm:"foreignKey:CosUUID;references:CosUUID"`
}

// OwnCostume 保有衣装テーブル
type OwnCostume struct {
	FriendUUID string `gorm:"column:FRIEND_UUID;type:CHAR(36);not null;foreignKey:FriendUUID;references:FriendUUID"`
	CosUUID    string `gorm:"column:COS_UUID;type:CHAR(36);not null;foreignKey:CosUUID;references:CosUUID"`
}

// QuestHistory クエスト履歴テーブル
type QuestHistory struct {
	QuestUUID  string    `gorm:"primaryKey;column:QUEST_UUID;type:CHAR(36);not null"`
	FriendUUID string    `gorm:"column:FRIEND_UUID;type:CHAR(36);not null"`
	StoreName  string    `gorm:"column:STORE_NAME;type:VARCHAR(30);not null"`
	StoreLoca  string    `gorm:"column:STORE_LOCA;type:VARCHAR(50);not null"`
	StoType    int       `gorm:"column:STO_TYPE;type:INT;not null"`
	Possible   int       `gorm:"column:POSSIBLE;type:INT;not null;default:1"`
	CreateAt   time.Time `gorm:"column:CREATE_AT;type:timestamp;not null"`
}

type QuestCheck struct {
    UserUUID   string    `gorm:"primaryKey;column:USER_UUID;type:CHAR(36);not null"` // ユーザー固有識別子
    FriendUUID string    `gorm:"column:FRIEND_UUID;type:CHAR(36);not null"`          // フレンド固有識別子
    CreateAt   time.Time `gorm:"column:create_at;type:timestamp;not null"`            // 達成日時
}

// StoreType 店舗種別テーブル
type StoreType struct {
	StoType  int    `gorm:"primaryKey;column:STO_TYPE;type:INT;not null"`
	TypeName string `gorm:"column:TYPE_NAME;type:VARCHAR(50);not null"`
}

// MeetHistory 遊んだ日履歴テーブル
type MeetHistory struct {
	MeetUUID   string    `gorm:"primaryKey;column:MEET_UUID;type:CHAR(36);not null"`
	FriendUUID string    `gorm:"column:FRIEND_UUID;type:CHAR(36);not null"`
	MeetAt     time.Time `gorm:"column:MEET_AT;type:DATE;not null"`
	Pictures   []Picture `gorm:"foreignKey:MeetUUID;references:MeetUUID"`
}

// Picture 写真テーブル
type Picture struct {
	PicUUID  string    `gorm:"primaryKey;column:PIC_UUID;type:CHAR(36);not null"`
	MeetUUID string    `gorm:"column:MEET_UUID;type:CHAR(36);not null"`
	PicURL   string    `gorm:"column:PIC_URL;type:VARCHAR(50);not null"`
	CreateAt time.Time `gorm:"column:CREATE_AT;type:timestamp;not null"`
}

//ここまでデータベース


