package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FollowerRepository interface {
	FindByUserID(db *gorm.DB, userID int) (followers []domain.Followers, err error)
	Create(db *gorm.DB, follower domain.Followers) (newFollower domain.Followers, err error)
}
