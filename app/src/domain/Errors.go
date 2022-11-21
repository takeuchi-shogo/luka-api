package domain

var (
	// User
	ErrGetUserAccount    = "ユーザー情報の取得に失敗しました"
	ErrCreateUserAccount = "ユーザーアカウントの作成に失敗しました"
	ErrUpdateUserAccount = "ユーザー情報の編集に失敗しました"
	ErrDeleteUserAccount = "ユーザー情報の削除に失敗しました"

	ErrUserNotFound = "ユーザーが見つかりません"

	ExistUserScreenName = "このログインIDは既に使用されています"
	// Token
	ErrSignIn          = "ログインID、メールアドレス、パスワードのいずれかに間違いがあります"
	ErrCreateUserToken = "トークンの発行に失敗しました"
	ErrAuthorization   = "ログインに失敗しました"

	ErrTokenExpire        = "承認に失敗しました"
	ErrRefreshTokenExpire = "トークンの有効期限が切れています"

	// Thread
	ErrCreateThread   = "スレッドの作成に失敗しました"
	ErrSaveThread     = "スレッドの編集に失敗しました"
	ErrDeleteThread   = "スレッドの削除に失敗しました"
	ErrThreadNotFound = "スレッドが見つかりません"

	// Comment
	ErrCreateComment = "コメントの作成に失敗しました"

	ErrCommentNotFound = "コメントがありません"

	// Follower
	ErrCreateFollower = "フォロワーの作成に失敗しました"

	ErrFollowerNotFound = "フォローしているユーザーはいません"

	// Favorites
	ErrFavoriteCommentNotFound = "まだこのコメントにいいねはありません"
	ErrFavoriteThreadNotFound  = "まだこのスレッドにいいねはありません"
)
