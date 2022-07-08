package domain

type Threads struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}
