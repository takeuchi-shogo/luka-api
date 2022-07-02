package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserRepository interface {
	FindByScreenName(db *gorm.DB, userInfo string) (user domain.Users, err error)
}
