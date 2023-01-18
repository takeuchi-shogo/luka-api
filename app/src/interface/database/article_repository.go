package database

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ArticleRepository struct{}

func (r *ArticleRepository) Find(db *gorm.DB) (threads []domain.Articles, err error) {
	threads = []domain.Articles{}
	db.Find(&threads)
	if len(threads) <= 0 {
		return []domain.Articles{}, err
	}
	return threads, nil
}

func (r *ArticleRepository) FindByID(db *gorm.DB, id int) (foundThread domain.Articles, err error) {

	foundThread = domain.Articles{}

	db.First(&foundThread, id)
	if foundThread.ID < 0 {
		return domain.Articles{}, errors.New("test error")
	}

	return foundThread, nil
}

func (r *ArticleRepository) FindByUserID(db *gorm.DB, userID int) (foundThreads []domain.Articles, err error) {
	return foundThreads, nil
}

func (r *ArticleRepository) CountByUserID(db *gorm.DB, userID int) (threadCnt int, err error) {
	threadCnt = 0
	db.Model(&domain.Articles{}).Where("user_id = ?", userID).Count(&threadCnt)
	return threadCnt, nil
}

func (r *ArticleRepository) Create(db *gorm.DB, article domain.Articles) (newThread domain.Articles, err error) {

	newThread = domain.Articles{}

	newThread.UserID = article.UserID
	newThread.Title = article.Title
	newThread.Description = article.Description

	currentTime := time.Now().Unix()
	newThread.CreatedAt = currentTime
	newThread.UpdatedAt = currentTime
	newThread.DeletedAt = nil

	if err := newThread.Validate(); err != nil {
		return domain.Articles{}, err
	}

	db.NewRecord(&newThread)
	err = db.Create(&newThread).Error

	return newThread, err
}

func (r *ArticleRepository) Save(db *gorm.DB, article domain.Articles) (updateThread domain.Articles, err error) {

	updateThread = domain.Articles{}

	updateThread.ID = article.ID
	updateThread.UserID = article.UserID
	updateThread.Title = article.Title
	updateThread.Description = article.Description
	updateThread.CreatedAt = article.CreatedAt
	updateThread.UpdatedAt = time.Now().Unix()
	updateThread.DeletedAt = nil

	if err := updateThread.Validate(); err != nil {
		return domain.Articles{}, err
	}

	if err := db.Save(&updateThread).Error; err != nil {
		return domain.Articles{}, err
	}

	return updateThread, nil
}

// 論理削除
func (r *ArticleRepository) Delete(db *gorm.DB, article domain.Articles) error {
	currentTime := time.Now().Unix()
	article.DeletedAt = &currentTime
	return db.Save(&article).Error
}
