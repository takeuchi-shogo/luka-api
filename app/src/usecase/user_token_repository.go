package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserTokenRepository interface {
	FindByToken(db *gorm.DB, token string) (foundToken domain.UserTokens, err error)
	Create(db *gorm.DB, token domain.UserTokens) (newToken domain.UserTokens, err error)
}
