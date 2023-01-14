package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FollowingRepository interface {
	FindByUserID(db *gorm.DB, userID int) (followings []domain.Followings, err error)
	Create(db *gorm.DB, following domain.Followings) (newFollowing domain.Followings, err error)
	Delete(db *gorm.DB, toUserID int) error
}
