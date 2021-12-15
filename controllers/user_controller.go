package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"irisStudy/repositories"
	"net/http"
)

type createUserRequest struct {
	Email string `json:"email" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	HashPassword string `json:"hash_password" binding:"required"`
}

func (controller *APIController) CreatUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.InsertUserParams{
		Email: req.Email,
		UserName: req.UserName,
		HashPassword: req.HashPassword,
	}

	user, err := controller.store.InsertUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResponse := map[string]interface{}{
		"email": user.Email,
		"user_name": user.UserName,
	}

	ctx.JSON(http.StatusOK, userResponse)
}

type deleteUserRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (controller *APIController) DeleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.DeleteUserParams{
		ID: req.ID,
	}

	if ok := controller.store.DeleteUser(ctx, arg); !ok {
		err := errors.New("failed to delete user, err")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete success"})
}


