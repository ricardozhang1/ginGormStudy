package datamodels

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserID int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	OrderStatus int `json:"order_status"`
}

const (
	OrderWait = iota
	OrderSuccess  //1
	OrderFailed //2
)
