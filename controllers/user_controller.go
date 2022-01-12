package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"irisStudy/repositories"
	"irisStudy/utils"
	"net/http"
	"strconv"
	"time"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"user_name" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

func (controller *APIController) CreatUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := repositories.InsertUserParams{
		Email:        req.Email,
		UserName:     req.UserName,
		HashPassword: hashPassword,
	}

	user, err := controller.store.InsertUser(ctx, arg)
	if err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createUserResponse{
		Email:    user.Email,
		UserName: user.UserName,
	}

	ctx.JSON(http.StatusOK, resp)
}

type deleteUserRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (controller *APIController) DeleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.DeleteUserParams{
		ID: req.ID,
	}

	if ok := controller.store.DeleteUser(ctx, arg); !ok {
		err := errors.New("failed to delete user, err")
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete success"})
}

type logUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type logUserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	AccessToken string `json:"access_token"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (controller *APIController) LogUser(ctx *gin.Context) {
	var req logUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.GetUserParams{
		Email: req.Email,
	}

	// 数据库中查询
	user, err := controller.store.GetUser(ctx, arg)
	if err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// 密码的验证 加密处理
	if err = utils.CheckPassword(req.Password, user.HashPassword); err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := controller.tokenMaker.CreateToken(user.UserName, strconv.Itoa(int(user.ID)), controller.config.AccessTokenDuration)
	if err != nil {
		controller.logger.ERROR(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := logUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		UserName:  user.UserName,
		AccessToken: accessToken,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}
