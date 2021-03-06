package domain

import (
	"crypto/sha256"
	"fmt"
)

type Users struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	ScreenName  string `json:"screenName"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Prefecture  string `json:"prefecture"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

type UserForPatch struct {
	ID          int
	DisplayName string
	ScreenName  string
	Password    string
	Email       string
	Age         int
	Gender      string
	Prefecture  string
}

// パスワードのハッシュ化
func (u *Users) GetPassword(password string) string {
	var data [sha256.Size]byte
	data = sha256.Sum256(([]byte(password)))
	return fmt.Sprintf("%x", data)
}
