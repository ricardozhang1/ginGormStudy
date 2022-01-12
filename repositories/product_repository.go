package repositories

import (
	"context"
	"irisStudy/datamodels"
)

type CreateProductParams struct {
	ProductName string `json:"product_name"`
	ProductNum int64 `json:"product_num"`
	ProductImage string `json:"product_image"`
	ProductUrl string `json:"product_url"`
}

func (sqlStore *SQLStore) CreateProduct(ctx context.Context, arg CreateProductParams) bool {
	var product datamodels.Product
	product.ProductName = arg.ProductName
	product.ProductNum = arg.ProductNum
	product.ProductImage = arg.ProductImage
	product.ProductUrl = arg.ProductUrl

	if err := sqlStore.db.Create(&product).Error; err != nil {
		return false
	}
	return true
}

type DeleteProductParams struct {
	ID int64 `json:"id"`
}

func (sqlStore *SQLStore) DeleteProduct(ctx context.Context, arg DeleteProductParams) bool {
	var product datamodels.Product
	product.ID = arg.ID

	if sqlStore.db.Delete(&product).Error != nil {
		return false
	}
	return true
}

type GetProductParams struct {
	ID int64 `json:"id"`
}

func (sqlStore *SQLStore) GetProduct(ctx context.Context, arg GetProductParams) (product datamodels.Product, err error) {
	if err = sqlStore.db.First(&product, "id = ?", arg.ID).Error; err != nil {
		return datamodels.Product{}, err
	}
	return
}

type UpdateProductParams struct {
	ID int64 `json:"id"`
	ProductName string `json:"product_name"`
	ProductNum int64 `json:"product_num"`
	ProductImage string `json:"product_image"`
	ProductUrl string `json:"product_url"`
}

func (sqlStore *SQLStore) UpdateProduct(ctx context.Context, arg UpdateProductParams) bool {
	var product datamodels.Product
	product.ProductUrl = arg.ProductUrl
	product.ProductName = arg.ProductName
	product.ProductNum = arg.ProductNum
	product.ProductImage = arg.ProductImage

	if err := sqlStore.db.Where("id = ?").Update(&product); err != nil {
		return false
	}

	return true
}



