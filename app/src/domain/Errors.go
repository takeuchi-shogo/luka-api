package domain

const (
	// User
	GetUserAccountError    = "ユーザー情報の取得に失敗しました。"
	CreateUserAccountError = "ユーザーアカウントの作成に失敗しました。"
	UpdateUserAccountError = "ユーザー情報の編集に失敗しました。"
	DeleteUserAccountError = "ユーザーアカウントの削除に失敗しました。"
	// Token
	SignInError          = "ログインID、メールアドレス、パスワードのいずれかに間違いがあります。"
	CreateUserTokenError = "トークンの発行に失敗しました。"
)
