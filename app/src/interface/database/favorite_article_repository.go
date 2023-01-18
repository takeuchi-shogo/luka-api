package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteArticleRepository struct{}

func (r *FavoriteArticleRepository) FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteArticles, err error) {

	favorites = []domain.FavoriteArticles{}

	db.Where("user_id = ?", userID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteArticles{}, errors.New(domain.ErrFavoriteCommentNotFound)
	}

	return favorites, nil
}

func (r FavoriteArticleRepository) FindByArticleID(db *gorm.DB, articleID int) (favorites []domain.FavoriteArticles, err error) {

	favorites = []domain.FavoriteArticles{}

	db.Where("article_id = ?", articleID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteArticles{}, errors.New(domain.ErrFavoriteArticleNotFound)
	}

	return favorites, nil
}
