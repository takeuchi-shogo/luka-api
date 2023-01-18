package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type CommentRepository interface {
	FindByArticleID(db *gorm.DB, articleID int) (foundComments []domain.Comments, err error)
	Create(db *gorm.DB, comment domain.Comments) (newComment domain.Comments, err error)
}
