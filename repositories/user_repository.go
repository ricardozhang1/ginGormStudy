package repositories

import (
	"context"
	"fmt"
	"irisStudy/datamodels"
)

type InsertUserParams struct {
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
}

func (sqlStore *SQLStore) InsertUser(ctx context.Context, arg InsertUserParams) (datamodels.User, error) {
	fmt.Println("执行User插入数据库操作。。。")
	data := datamodels.User{
		Email:        arg.Email,
		UserName:     arg.UserName,
		HashPassword: arg.HashPassword,
	}
	sqlStore.db.Create(&data)
	return datamodels.User{}, nil
}

type GetUserParams struct {
	Email string `json:"email"`
}

func (sqlStore *SQLStore) GetUser(ctx context.Context, arg GetUserParams) (user datamodels.User, err error) {
	fmt.Println("执行User查询数据库操作。。。")
	sqlStore.db.Where("email = ?", arg.Email).First(&user)
	return
}

type ListUsersParams struct {
	ID     int64 `json:"id"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
}

func (sqlStore *SQLStore) ListUsers(ctx context.Context, arg ListUsersParams) ([]datamodels.User, error) {
	return []datamodels.User{}, nil
}

type DeleteUserParams struct {
	ID int64 `json:"id"`
}

func (sqlStore *SQLStore) DeleteUser(ctx context.Context, arg DeleteUserParams) bool {
	fmt.Println("执行User删除数据库操作。。。")
	data := datamodels.User{
		ID: arg.ID,
	}
	sqlStore.db.Delete(data)
	return true
}

type UpdateUser struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	HashPassword string `json:"hash_password"`
}

func (sqlStore *SQLStore) UpdateUser(ctx context.Context, arg UpdateUser) {
	data := datamodels.User{
		ID:           arg.ID,
		Email:        arg.Email,
		UserName:     arg.UserName,
		HashPassword: arg.HashPassword,
	}
	sqlStore.db.Model(&datamodels.User{}).Update(data)
}
