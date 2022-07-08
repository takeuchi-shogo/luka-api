package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ThreadRepository struct{}

func (r *ThreadRepository) Create(db *gorm.DB, thread domain.Threads) (newThread domain.Threads, err error) {

	newThread = domain.Threads{}

	newThread.UserID = thread.UserID
	newThread.Title = thread.Title
	newThread.Description = thread.Description

	currentTime := time.Now().Unix()
	newThread.CreatedAt = currentTime
	newThread.UpdatedAt = currentTime

	db.NewRecord(&newThread)
	err = db.Create(&newThread).Error
	if err != nil {
		return domain.Threads{}, err
	}

	return newThread, nil
}
