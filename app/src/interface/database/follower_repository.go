package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FollowerRepository struct{}

func (r *FollowerRepository) FindByUserID(db *gorm.DB, userID int) (followers []domain.Followers, err error) {
	followers = []domain.Followers{}
	db.Where("user_id = ?", userID).Find(&followers)
	if len(followers) <= 0 {
		return []domain.Followers{}, err
	}
	return followers, nil
}

func (r *FollowerRepository) Create(db *gorm.DB, follower domain.Followers) (newFollower domain.Followers, err error) {

	newFollower = domain.Followers{}

	newFollower.UserID = follower.UserID
	newFollower.ToUserID = follower.ToUserID
	newFollower.CreatedAt = time.Now().Unix()

	db.NewRecord(&newFollower)
	err = db.Create(&newFollower).Error

	return newFollower, err
}
