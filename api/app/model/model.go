package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User data structure
type User struct {
	Id       int     `json:"id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Location float64 `json:"location" binding:"required"`
	Gender   string  `json:"gender,omitempty"`
	Email    string  `json:"email,omitempty"`
}

//Like data structure
type Like struct {
	Id           int `json:"id"`
	Who_likes    int `json:"who_likes"`
	Who_is_liked int `json:"who_is_liked"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Like{})
	return db
}
