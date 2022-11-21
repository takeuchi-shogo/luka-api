package domain

type FavoriteThreads struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	ThreadID int `json:"threadId"`

	CreatedAt int64 `json:"createdAt"`
}
