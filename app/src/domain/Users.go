package domain

type Users struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	ScreenName  string `json:"screenName"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Prefecture  string `json:"prefecture"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}
