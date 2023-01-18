package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ArticleRepository interface {
	Find(db *gorm.DB) (articles []domain.Articles, err error) // Test
	FindByID(db *gorm.DB, id int) (article domain.Articles, err error)
	CountByUserID(db *gorm.DB, userID int) (articleCnt int, err error)
	Create(db *gorm.DB, article domain.Articles) (newThead domain.Articles, err error)
	Save(db *gorm.DB, article domain.Articles) (updateArticle domain.Articles, err error)
	Delete(db *gorm.DB, article domain.Articles) error
}
