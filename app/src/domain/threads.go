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

type ThreadsForGet struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`

	User     Users      `json:"user"`
	Comments []Comments `json:"comments"`
}

type ThreadsForPatch struct {
	ID          int
	UserID      int
	Title       string
	Description string
}

func (t *Threads) BuildForGet() ThreadsForGet {
	thread := ThreadsForGet{}

	thread.ID = t.ID
	thread.UserID = t.UserID
	thread.Title = t.Title
	thread.Description = t.Description
	thread.CreatedAt = t.CreatedAt

	return thread
}
