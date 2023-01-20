package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ArticleRepository struct{}

func (r *ArticleRepository) Find(db *gorm.DB) (articles []domain.Articles, err error) {
	articles = []domain.Articles{}
	db.Find(&articles)
	if len(articles) <= 0 {
		return []domain.Articles{}, err
	}
	return articles, nil
}

func (r *ArticleRepository) FindByID(db *gorm.DB, id int) (foundArticle domain.Articles, err error) {

	foundArticle = domain.Articles{}

	db.First(&foundArticle, id)
	if foundArticle.ID <= 0 {
		return domain.Articles{}, fmt.Errorf("article is not found: %v", id)
	}

	return foundArticle, nil
}

func (r *ArticleRepository) FindByUserID(db *gorm.DB, userID int) (foundArticles []domain.Articles, err error) {
	return foundArticles, nil
}

func (r *ArticleRepository) CountByUserID(db *gorm.DB, userID int) (articleCnt int, err error) {
	articleCnt = 0
	db.Model(&domain.Articles{}).Where("user_id = ?", userID).Count(&articleCnt)
	return articleCnt, nil
}

func (r *ArticleRepository) Create(db *gorm.DB, article domain.Articles) (newArticle domain.Articles, err error) {

	newArticle = domain.Articles{}

	newArticle.UserID = article.UserID
	newArticle.Title = article.Title
	newArticle.Description = article.Description

	currentTime := time.Now().Unix()
	newArticle.CreatedAt = currentTime
	newArticle.UpdatedAt = currentTime
	newArticle.DeletedAt = nil

	if err := newArticle.Validate(); err != nil {
		return domain.Articles{}, err
	}

	db.NewRecord(&newArticle)
	err = db.Create(&newArticle).Error

	return newArticle, err
}

func (r *ArticleRepository) Save(db *gorm.DB, article domain.Articles) (updateArticle domain.Articles, err error) {

	updateArticle = domain.Articles{}

	updateArticle.ID = article.ID
	updateArticle.UserID = article.UserID
	updateArticle.Title = article.Title
	updateArticle.Description = article.Description
	updateArticle.CreatedAt = article.CreatedAt
	updateArticle.UpdatedAt = time.Now().Unix()
	updateArticle.DeletedAt = nil

	fmt.Println(updateArticle)

	if err := updateArticle.Validate(); err != nil {
		return domain.Articles{}, err
	}

	if err := db.Save(&updateArticle).Error; err != nil {
		return domain.Articles{}, err
	}

	return updateArticle, nil
}

// 論理削除
func (r *ArticleRepository) Delete(db *gorm.DB, article domain.Articles) error {
	currentTime := time.Now().Unix()
	article.DeletedAt = &currentTime
	return db.Save(&article).Error
}
