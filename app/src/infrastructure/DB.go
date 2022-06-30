package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(c *Config) *DB {
	db := &DB{}

	db.Connection = db.connect(
		c.DB.Production.Host,
		c.DB.Production.UserName,
		c.DB.Production.Password,
		c.DB.Production.DBName,
	)
	return db
}

// DBに接続
func (db *DB) connect(host string, username string, password string, dbName string) *gorm.DB {
	// https://github.com/go-sql-driver/mysql#examples
	connection, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName))
	if err != nil {
		panic(err.Error())
	}
	return connection
}
