package datamodels

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	ProductNum int64 `json:"product_num"`
	ProductImage string `json:"product_image"`
	ProductUrl string `json:"product_url"`
}



