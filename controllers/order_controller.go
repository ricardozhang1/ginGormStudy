package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"irisStudy/repositories"
	"net/http"
)

type createOrderRequest struct {
	UserID      int64 `json:"user_id" binding:"required"`
	ProductID   int64 `json:"product_id" binding:"required"`
	OrderStatus int   `json:"order_status" binding:"required"`
}

func (controller *APIController) CreatOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.InsertOrderParams{
		UserID: req.UserID,
		ProductID: req.ProductID,
		OrderStatus: req.OrderStatus,
	}
	fmt.Println(arg)


	if ok := controller.store.InsertOrder(ctx, arg); !ok {
		err := errors.New("failed to create order")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}


type getOrderByUserIDRequest struct {
	UserID int64 `uri:"user_id" binding:"required"`
}

func (controller *APIController) GetOrderByUser(ctx *gin.Context) {
	var req getOrderByUserIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.GetOrderByUserParams{
		UserID: req.UserID,
	}

	orders, err := controller.store.GetByUserOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

type getOrderByProductIDRequest struct {
	ProductID int64 `uri:"product_id" binding:"required"`
}

func (controller *APIController) GetOrderByProduct(ctx *gin.Context) {
	var req getOrderByProductIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.GetOrderByProductParams{
		ProductID: req.ProductID,
	}

	orders, err := controller.store.GetByProductOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

