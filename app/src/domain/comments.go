package domain

type Comments struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	ToUserID  int    `json:"toUserID"`
	ArticleID int    `json:"articleId"`
	Content   string `json:"content"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}
