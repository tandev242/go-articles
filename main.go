package main

import (
	"go-training/go-articles/common"
	"go-training/go-articles/users"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {
	db := common.Init()
	Migrate(db)
	// close db when app crashed or quit
	defer db.Close()
	router := gin.Default()
	router.Run("App is running on port 8080")
}
