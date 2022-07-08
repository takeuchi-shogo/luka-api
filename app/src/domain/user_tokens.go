package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type UserTokens struct {
	ID                    int    `json:"id"`
	UserID                int    `json:"userId"`
	Token                 string `json:"token"`
	TokenExpiredAt        int64  `json:"tokenExpiredAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiredAt int64  `json:"refreshTokenExpiredAt"`

	CreatedAt int64 `json:"createdAt"`
}

func (u *UserTokens) GetToken() string {

	source := rand.NewSource(time.Now().UnixNano())

	maxRange := 50
	minRange := 40

	str := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	cnt := len(str)

	var token string

	for i := 0; i < maxRange; i++ {
		if minRange <= i {
			if rand.New(source).Intn(maxRange-minRange) == 0 {
				break
			}
		}
		token = token + fmt.Sprintf("%c", str[rand.New(source).Intn(cnt)])
	}

	return token
}
