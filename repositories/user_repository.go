package repositories

import (
	"context"
	"fmt"
	"irisStudy/datamodels"
)

type InsertUserParams struct {
	Email string `json:"email"`
	UserName string `json:"user_name"`
	HashPassword string `json:"hash_password"`
}

func (sqlStore *SQLStore) InsertUser(ctx context.Context, arg InsertUserParams) (datamodels.User, error) {
	// 进行数据库迁移
	sqlStore.db.AutoMigrate(&datamodels.User{})

	fmt.Println("执行User插入数据库操作。。。")
	return datamodels.User{}, nil
}


type DeleteUserParams struct {
	ID int64 `json:"id"`
}

func (sqlStore *SQLStore) DeleteUser(ctx context.Context, arg DeleteUserParams) bool {
	fmt.Println("执行User删除数据库操作。。。")
	return true
}



