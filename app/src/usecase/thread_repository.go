package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type ThreadRepository interface {
	Create(db *gorm.DB, thread domain.Threads) (newThead domain.Threads, err error)
}
