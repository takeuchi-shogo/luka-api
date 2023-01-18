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

	// Article
	ErrCreateArticle   = "スレッドの作成に失敗しました"
	ErrSaveArticle     = "スレッドの編集に失敗しました"
	ErrDeleteArticle   = "スレッドの削除に失敗しました"
	ErrArticleNotFound = "スレッドが見つかりません"

	// Comment
	ErrCreateComment = "コメントの作成に失敗しました"

	ErrCommentNotFound = "コメントがありません"

	// Follower
	ErrCreateFollower = "フォロワーの作成に失敗しました"

	ErrFollowerNotFound = "フォローされているユーザーはいません"

	// Following
	ErrFollowingNotFound = "フォローしているユーザーはいません"

	// Favorites
	ErrFavoriteCommentNotFound = "まだこのコメントにいいねはありません"
	ErrFavoriteArticleNotFound = "まだこのスレッドにいいねはありません"
)
