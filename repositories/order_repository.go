package repositories

import (
	"context"
	"irisStudy/datamodels"
)


type InsertOrderParams struct {
	UserID      int64 `json:"user_id"`
	ProductID   int64 `json:"product_id"`
	OrderStatus int   `json:"order_status"`
}

// InsertOrder 创建订单
func (sqlStore *SQLStore) InsertOrder(ctx context.Context, arg InsertOrderParams) bool {
	var order datamodels.Order
	order.UserID = arg.UserID
	order.ProductID = arg.ProductID
	order.OrderStatus = arg.OrderStatus

	if err := sqlStore.db.Create(&order).Error; err != nil {
		return false
	}
	return true
}

type GetOrderByUserParams struct {
	UserID int64 `json:"user_id"`
}

// GetByUserOrder 根据UserID查询订单
func (sqlStore *SQLStore) GetByUserOrder(ctx context.Context, arg GetOrderByUserParams) (order []datamodels.Order, err error) {
	if err = sqlStore.db.Find(&order, "user_id = ?", arg.UserID).Error; err != nil {
		return
	}
	return
}

type GetOrderByProductParams struct {
	ProductID int64 `json:"product_id"`
}

// GetByProductOrder 根据productID查询订单
func (sqlStore *SQLStore) GetByProductOrder(ctx context.Context, arg GetOrderByProductParams) (order []datamodels.Order, err error) {
	if err = sqlStore.db.Find(&order, "product_id = ?", arg.ProductID).Error; err != nil {
		return
	}
	return
}





