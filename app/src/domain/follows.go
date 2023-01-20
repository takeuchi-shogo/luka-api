package domain

type Follows struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	ToUserID int `json:"toUserId"`

	CreatedAt int64
}
