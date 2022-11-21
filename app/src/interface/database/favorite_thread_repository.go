package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteThreadRepository struct{}

func (r *FavoriteThreadRepository) FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteThreads, err error) {

	favorites = []domain.FavoriteThreads{}

	db.Where("user_id = ?", userID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteThreads{}, errors.New(domain.ErrFavoriteCommentNotFound)
	}

	return favorites, nil
}

func (r FavoriteThreadRepository) FindByThreadID(db *gorm.DB, threadID int) (favorites []domain.FavoriteThreads, err error) {

	favorites = []domain.FavoriteThreads{}

	db.Where("thread_id = ?", threadID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteThreads{}, errors.New(domain.ErrFavoriteThreadNotFound)
	}

	return favorites, nil
}
