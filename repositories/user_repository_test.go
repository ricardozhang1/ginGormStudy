package repositories

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertUser(t *testing.T) {
	arg :=  InsertUserParams{
		Email: "73298428@gmail.com",
		UserName: "lisi",
		HashPassword: "2378628e723gd8326tde76g",
	}
	_, _ = testStore.InsertUser(context.Background(), arg)
	//require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	arg := GetUserParams{
		Email: "987657687@gmail.com",
	}
	user, _ := testStore.GetUser(context.Background(), arg)
	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	arg := UpdateUser{
		ID: 2,
		Email: "987657687@gmail.com",
		UserName: "lisi",
		HashPassword: "ds76d87scys8",
	}
	testStore.UpdateUser(context.Background(), arg)
}




