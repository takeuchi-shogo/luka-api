package domain

var (
	// User
	GetUserAccountError    = "ユーザー情報の取得に失敗しました。"
	CreateUserAccountError = "ユーザーアカウントの作成に失敗しました。"
	UpdateUserAccountError = "ユーザー情報の編集に失敗しました。"
	DeleteUserAccountError = "ユーザーアカウントの削除に失敗しました。"

	ErrUserNotFound = "ユーザーが見つかりません。"

	ExistUserScreenName = "このログインIDは既に使用されています。"
	// Token
	SignInError          = "ログインID、メールアドレス、パスワードのいずれかに間違いがあります。"
	CreateUserTokenError = "トークンの発行に失敗しました。"
	ErrAuthorization     = "ログインに失敗しました。"

	ErrTokenExpire        = "承認に失敗しました。"
	ErrRefreshTokenExpire = "トークンの有効期限が切れています。"

	// Thread
	ErrCreateThread = "スレッドの作成に失敗しました。"
)
