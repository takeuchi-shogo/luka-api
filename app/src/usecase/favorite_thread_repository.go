package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteThreadRepository interface {
	FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteThreads, err error)
	FindByThreadID(db *gorm.DB, threadID int) (favorites []domain.FavoriteThreads, err error)
}
