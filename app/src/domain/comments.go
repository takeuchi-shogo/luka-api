package domain

type Comments struct {
	ID       int    `json:"id"`
	ThreadID int    `json:"threadId"`
	UserID   int    `json:"userId"`
	Content  string `json:"content"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}
