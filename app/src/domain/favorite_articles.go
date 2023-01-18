package domain

type FavoriteArticles struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	ArticleID int `json:"articleId"`

	CreatedAt int64 `json:"createdAt"`
}
