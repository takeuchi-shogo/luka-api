package database

import (
	"time"

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

func (r *UserTokenRepository) Create(db *gorm.DB, userToken domain.UserTokens) (newToken domain.UserTokens, err error) {

	newToken = domain.UserTokens{}

	var token string

	for {
		token = newToken.GetToken()
		u, err := r.FindByToken(db, token)
		if err != nil {
			break
		}
		if u.TokenExpiredAt < time.Now().Unix() {
			break
		}
	}

	newToken.UserID = userToken.UserID
	newToken.Token = token
	newToken.RefreshToken = newToken.GetToken()
	newToken.CreatedAt = time.Now().Unix()
	newToken.SetExpireAt()

	db.NewRecord(newToken)
	err = db.Create(&newToken).Error

	return newToken, err
}
