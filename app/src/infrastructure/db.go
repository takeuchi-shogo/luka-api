package infrastructure

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
// docker-composeしたときDBよりもアプリが早く立ち上がってしまうので10秒間待機する
func (db *DB) connect(host string, username string, password string, dbName string) *gorm.DB {
	count := 0
	// https://github.com/go-sql-driver/mysql#examples
	connection, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName))
	if err != nil {
		for {
			if err == nil {
				break
			}
			fmt.Print(".\n")
			time.Sleep(time.Second)
			count++
			if count > 10 {
				fmt.Print("database connection failed\n")
				panic(err.Error())
			}
			connection, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName))

		}
	}

	fmt.Print("database connection success\n")

	return connection
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
