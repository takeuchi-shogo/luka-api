package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteArticleRepository interface {
	FindByID(db *gorm.DB, id int) (domain.FavoriteArticles, error)
	FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteArticles, err error)
	FindByArticleID(db *gorm.DB, articleID int) (favorites []domain.FavoriteArticles, err error)
	Create(db *gorm.DB, favorite domain.FavoriteArticles) (newFavorite domain.FavoriteArticles, err error)
	Delete(db *gorm.DB, id int) error
}
