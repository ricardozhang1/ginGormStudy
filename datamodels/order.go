package datamodels

import "time"

type Order struct {
	ID int64 `json:"id" gorm:"primary_key,auto_increment"`
	UserID int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	OrderStatus int `json:"order_status"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at" gorm:"index"`
}

const (
	OrderWait = iota
	OrderSuccess  //1
	OrderFailed //2
)
