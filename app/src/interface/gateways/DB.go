package gateways

import "github.com/jinzhu/gorm"

type DB interface {
	Connect() *gorm.DB
}
