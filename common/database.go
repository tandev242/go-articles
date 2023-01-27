package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save to `Database` struct above
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../../gorm.db")
	if err != nil {
		fmt.Println("Can not connect to db: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
