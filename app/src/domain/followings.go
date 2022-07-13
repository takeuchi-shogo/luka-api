package domain

type Followings struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	ToUserID int `json:"toUserId"`

	CreatedAt int64
}
