package datamodels

import "time"

type Product struct {
	ID int64 `json:"id" gorm:"primary_key,auto_increment"`
	ProductName string `json:"product_name"`
	ProductNum int64 `json:"product_num"`
	ProductImage string `json:"product_image"`
	ProductUrl string `json:"product_url"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at" gorm:"index"`
}



