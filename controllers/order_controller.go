package controllers

import (
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

	order, err := controller.store.InsertOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}
