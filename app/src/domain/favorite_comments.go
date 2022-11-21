package domain

type FavoriteComments struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	CommentID int `json:"commentId"`

	CreatedAt int64 `json:"createdAt"`
}
