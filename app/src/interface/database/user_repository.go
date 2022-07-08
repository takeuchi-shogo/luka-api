package database

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserRepository struct{}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	return user, nil
}

func (repo *UserRepository) FindByScreenName(db *gorm.DB, screenName string) (user domain.Users, err error) {
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, user domain.Users) (newUser domain.Users, err error) {
	return newUser, nil
}
