package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserTokenRepository interface {
	Create(db *gorm.DB, token domain.UserTokens) (newToken domain.UserTokens, err error)
}
