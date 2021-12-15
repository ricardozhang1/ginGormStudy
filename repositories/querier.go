package repositories

import (
	"context"
	"irisStudy/datamodels"
)

type Querier interface {
	InsertOrder(context.Context, InsertOrderParams) (datamodels.Order, error)
	InsertUser(context.Context, InsertUserParams) (datamodels.User, error)
	DeleteUser(context.Context, DeleteUserParams) bool
}


