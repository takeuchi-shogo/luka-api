package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteCommentRepository struct{}

func (r *FavoriteCommentRepository) FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteComments, err error) {

	favorites = []domain.FavoriteComments{}

	db.Where("user_id = ?", userID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteComments{}, errors.New(domain.ErrFavoriteCommentNotFound)
	}

	return favorites, nil
}

func (r *FavoriteCommentRepository) FindByCommentID(db *gorm.DB, commentID int) (favorites []domain.FavoriteComments, err error) {

	favorites = []domain.FavoriteComments{}

	db.Where("comment_id = ?", commentID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteComments{}, errors.New(domain.ErrFavoriteCommentNotFound)
	}

	return favorites, nil
}
