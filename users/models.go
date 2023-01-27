package users

import (
	"go-training/go-articles/common"

	"github.com/jinzhu/gorm"
)

// Creating model map to DB schema and validating
// More detail at here: http://jinzhu.me/gorm/models.html#model-definition
// HINT: `*` is used to split null or "",
// exg: should *string instead of string.
type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

type FollowModel struct {
	gorm.Model
	Following   UserModel
	FollowingID uint
	FollowedBy  UserModel
	FollowedID  uint
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&FollowModel{})
}
