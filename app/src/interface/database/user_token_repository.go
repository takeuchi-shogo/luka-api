package database

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserTokenRepository struct{}

func (r *UserTokenRepository) FindByToken(db *gorm.DB, token string) (foundToken domain.UserTokens, err error) {

	foundToken = domain.UserTokens{}

	db.Where("token = ?", token).First(&foundToken)
	if foundToken.ID <= 0 {
		return domain.UserTokens{}, err
	}

	return foundToken, nil
}

func (repo *UserTokenRepository) Create(db *gorm.DB, token domain.UserTokens) (newToken domain.UserTokens, err error) {
	return newToken, nil
}
