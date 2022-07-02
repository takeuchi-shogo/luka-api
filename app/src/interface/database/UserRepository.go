package database

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserRepository struct{}

func (repo *UserRepository) FindByScreenName(db *gorm.DB, screenName string) (user domain.Users, err error) {
	return user, nil
}
