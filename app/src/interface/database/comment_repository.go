package database

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type CommentRepository struct{}

func (r *CommentRepository) FindByThreadID(db *gorm.DB, threadID int) (foundComments []domain.Comments, err error) {

	foundComments = []domain.Comments{}

	db.Where("thread_id = ?", threadID).Find(&foundComments)
	if len(foundComments) < 0 {
		return []domain.Comments{}, err
	}
	return foundComments, nil
}
