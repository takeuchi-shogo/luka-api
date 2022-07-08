package domain

type UserTokens struct {
	ID                    int    `json:"id"`
	UserID                int    `json:"userId"`
	Token                 string `json:"token"`
	TokenExpiredAt        int64  `json:"tokenExpiredAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiredAt int64  `json:"refreshTokenExpiredAt"`

	CreatedAt int64 `json:"createdAt"`
}
