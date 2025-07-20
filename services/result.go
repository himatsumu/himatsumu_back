package services

type Result struct {
	Message string      `json:"Message"`
	Status  int         `json:"Status"`
	Data    interface{} `json:"Data"`
}

const (
	notaccept = 0 //未承認
	accept    = 1 //承認
	reject    = 2 //拒否
	cancel    = 3 //送信キャンセル
)

const (
	AlreadySent              = "すでにリクエストを送信しています"
	SameUser                 = "同一ユーザーです"
	AlreadyFriends           = "既にフレンドです"
	UserNotFound             = "ユーザーが見つかりませんでした"
	FriendRegistrationFailed = "フレンドを登録できませんでした"
	RequestRegistrationFailed= "リクエストを送信できませんでした"
	RequestNotFound          = "リクエストが存在しませんでした"
	UserInfoFailed           = "ユーザー情報取得に失敗しました"
	Incorrectrequesterror    = "フレンドリクエストが無効です"
	UserMismatchExisting     = "ユーザーが一致していません"
	CharacterNotRegistration = "キャラクターを生成できませんでした"
	CouldNotGenerateName     = "名前を生成できませんでした"
	QuestNotCompleted        = "クエストを達成できていません"
	EmptyUUID                = "UUIDが空です"
	EmptyInfo                = "必要な情報が空です"
	InvalidDateTime          = "日付の形式が不正です"
	AlreadyUUID              = "UUIDが重複しています"
	UnexpectedError          = "予期しないエラーが発生しました"
	NotRecommendQuest        = "クエスト生成中にエラーが発生しました"
	NotFriend                = "フレンドではないユーザーです"
	QuestNotFound            = "クエストが見つかりませんでした"

	InvalidRequestFormat = "リクエストの形式が不正です"
)
