package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type CommentRepository interface {
	FindByThreadID(db *gorm.DB, threadID int) (foundComments []domain.Comments, err error)
}
