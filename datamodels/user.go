package datamodels

import "time"

type User struct {
	ID int64 `json:"id" gorm:"primary_key,auto_increment"`
	Email string `json:"email" gorm:"unique_index"`
	UserName string `json:"user_name"`
	HashPassword string `json:"hash_password"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at" gorm:"index"`
}
