package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ThreadRepository interface {
	FindByID(db *gorm.DB, id int) (thread domain.Threads, err error)
	Create(db *gorm.DB, thread domain.Threads) (newThead domain.Threads, err error)
	Save(db *gorm.DB, thread domain.Threads) (updateThread domain.Threads, err error)
	Delete(db *gorm.DB, thread domain.Threads) error
}
