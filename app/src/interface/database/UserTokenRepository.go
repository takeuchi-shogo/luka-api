package database

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserTokenRepository struct{}

func (repo *UserTokenRepository) Create(db *gorm.DB, token domain.UserTokens) (newToken domain.UserTokens, err error) {
	return newToken, nil
}
