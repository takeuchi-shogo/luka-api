package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
	FindByScreenName(db *gorm.DB, userInfo string) (user domain.Users, err error)
	Create(db *gorm.DB, user domain.Users) (newUser domain.Users, err error)
}
