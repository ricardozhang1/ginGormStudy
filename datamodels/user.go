package datamodels

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `json:"email"`
	UserName string `json:"user_name"`
	HashPassword string `json:"hash_password"`
}
