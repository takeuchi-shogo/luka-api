package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type CommentRepository struct{}

func (r *CommentRepository) FindByID(db *gorm.DB, id int) (comment domain.Comments, err error) {
	comment = domain.Comments{}
	db.Find(id, &comment)
	if comment.ID <= 0 {
		return domain.Comments{}, err
	}
	return comment, nil
}

func (r *CommentRepository) FindByArticleID(db *gorm.DB, articleID int) (foundComments []domain.Comments, err error) {

	foundComments = []domain.Comments{}

	db.Where("article_id = ?", articleID).Find(&foundComments)
	if len(foundComments) < 0 {
		return []domain.Comments{}, err
	}
	return foundComments, nil
}

func (r *CommentRepository) Create(db *gorm.DB, comment domain.Comments) (newComment domain.Comments, err error) {

	newComment = domain.Comments{}

	newComment.ArticleID = comment.ArticleID
	newComment.UserID = comment.UserID
	newComment.Content = comment.Content

	currentTime := time.Now().Unix()
	newComment.CreatedAt = currentTime
	newComment.UpdatedAt = currentTime
	newComment.DeletedAt = nil

	db.NewRecord(&newComment)
	err = db.Create(&newComment).Error

	return newComment, err
}
