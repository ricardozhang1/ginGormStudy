package repositories

import (
	"context"
	"fmt"
	"irisStudy/datamodels"
)


type InsertOrderParams struct {
	UserID      int64 `json:"user_id"`
	ProductID   int64 `json:"product_id"`
	OrderStatus int   `json:"order_status"`
}

func (sqlStore *SQLStore) InsertOrder(ctx context.Context, arg InsertOrderParams) (datamodels.Order, error) {
	// 进行数据库迁移
	sqlStore.db.AutoMigrate(&datamodels.Order{})

	fmt.Println("执行Order插入数据库操作。。。")
	return datamodels.Order{}, nil
}
