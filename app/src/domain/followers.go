package domain

type Followers struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	ToUserID int `json:"toUserId"`

	CreatedAt int64
}
