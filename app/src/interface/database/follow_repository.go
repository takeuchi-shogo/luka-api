package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FollowRepository struct{}

func (r *FollowRepository) FindByUserID(db *gorm.DB, userID int) (followers []domain.Follows, err error) {
	followers = []domain.Follows{}
	db.Where("user_id = ?", userID).Find(&followers)
	if len(followers) <= 0 {
		return []domain.Follows{}, err
	}
	return followers, nil
}

func (r *FollowRepository) CountByUserID(db *gorm.DB, userID int) (followerCnt int, err error) {
	followerCnt = 0
	db.Model(&domain.Follows{}).Where("user_id = ?", userID).Count(&followerCnt)
	return followerCnt, nil
}

func (r *FollowRepository) CountByToUserID(db *gorm.DB, userID int) (followerCnt int, err error) {
	followerCnt = 0
	db.Model(&domain.Follows{}).Where("to_user_id = ?", userID).Count(&followerCnt)
	return followerCnt, nil
}

func (r *FollowRepository) Create(db *gorm.DB, follower domain.Follows) (newFollower domain.Follows, err error) {

	newFollower = domain.Follows{}

	newFollower.UserID = follower.UserID
	newFollower.ToUserID = follower.ToUserID
	newFollower.CreatedAt = time.Now().Unix()

	db.NewRecord(&newFollower)
	err = db.Create(&newFollower).Error

	return newFollower, err
}
