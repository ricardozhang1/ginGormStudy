package repositories

import (
	"context"
	"irisStudy/datamodels"
)

type Querier interface {
	// Order的数据库操作
	InsertOrder(ctx context.Context, arg InsertOrderParams) bool
	GetByUserOrder(ctx context.Context, arg GetOrderByUserParams) (order []datamodels.Order, err error)
	GetByProductOrder(ctx context.Context, arg GetOrderByProductParams) (order []datamodels.Order, err error)

	// User的数据库操作
	InsertUser(context.Context, InsertUserParams) (datamodels.User, error)
	GetUser(context.Context, GetUserParams) (datamodels.User, error)
	ListUsers(context.Context, ListUsersParams) ([]datamodels.User, error)
	DeleteUser(context.Context, DeleteUserParams) bool
	UpdateUser(ctx context.Context, arg UpdateUser)

	// Product的数据库操作
	GetProduct(ctx context.Context, arg GetProductParams) (product datamodels.Product, err error)
	DeleteProduct(ctx context.Context, arg DeleteProductParams) bool
	CreateProduct(ctx context.Context, arg CreateProductParams) bool
	UpdateProduct(ctx context.Context, arg UpdateProductParams) bool
}


