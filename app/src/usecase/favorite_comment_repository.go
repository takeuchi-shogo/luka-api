package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteCommentRepository interface {
	FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteComments, err error)
	FindByCommentID(db *gorm.DB, commentID int) (favorites []domain.FavoriteComments, err error)
}
