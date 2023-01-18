package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteArticleRepository interface {
	FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteArticles, err error)
	FindByArticleID(db *gorm.DB, articleID int) (favorites []domain.FavoriteArticles, err error)
}
