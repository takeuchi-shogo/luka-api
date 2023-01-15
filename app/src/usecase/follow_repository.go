package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FollowRepository interface {
	FindByUserID(db *gorm.DB, userID int) (followers []domain.Follows, err error)
	CountByUserID(db *gorm.DB, userID int) (followerCnt int, err error)
	CountByToUserID(db *gorm.DB, userID int) (followerCnt int, err error)
	Create(db *gorm.DB, follower domain.Follows) (newFollower domain.Follows, err error)
}
