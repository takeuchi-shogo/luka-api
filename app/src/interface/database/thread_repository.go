package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ThreadRepository struct{}

func (r *ThreadRepository) Find(db *gorm.DB) (threads []domain.Threads, err error) {
	threads = []domain.Threads{}
	db.Find(&threads)
	if len(threads) <= 0 {
		return []domain.Threads{}, err
	}
	return threads, nil
}

func (r *ThreadRepository) FindByID(db *gorm.DB, id int) (foundThread domain.Threads, err error) {

	foundThread = domain.Threads{}

	db.Where("id = ?", id).First(&foundThread)
	if foundThread.ID <= 0 {
		return domain.Threads{}, err
	}

	return foundThread, nil
}

func (r *ThreadRepository) FindByUserID(db *gorm.DB, userID int) (foundThreads []domain.Threads, err error) {
	return foundThreads, nil
}

func (r *ThreadRepository) Create(db *gorm.DB, thread domain.Threads) (newThread domain.Threads, err error) {

	newThread = domain.Threads{}

	newThread.UserID = thread.UserID
	newThread.Title = thread.Title
	newThread.Description = thread.Description

	currentTime := time.Now().Unix()
	newThread.CreatedAt = currentTime
	newThread.UpdatedAt = currentTime
	newThread.DeletedAt = nil

	db.NewRecord(&newThread)
	err = db.Create(&newThread).Error

	return newThread, err
}

func (r *ThreadRepository) Save(db *gorm.DB, thread domain.Threads) (updateThread domain.Threads, err error) {

	updateThread = domain.Threads{}

	updateThread.ID = thread.ID
	updateThread.UserID = thread.UserID
	updateThread.Title = thread.Title
	updateThread.Description = thread.Description
	updateThread.CreatedAt = thread.CreatedAt
	updateThread.UpdatedAt = time.Now().Unix()
	updateThread.DeletedAt = nil

	if err := db.Save(&updateThread).Error; err != nil {
		return domain.Threads{}, err
	}

	return updateThread, nil
}

func (r *ThreadRepository) Delete(db *gorm.DB, thread domain.Threads) error {
	return db.Delete(&thread).Error
}
