package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"irisStudy/datamodels"
	"irisStudy/repositories"
	"net/http"
)

type createProductRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	ProductNum int64 `json:"product_num" binding:"required"`
	ProductImage string `json:"product_image"`
	ProductUrl string `json:"product_url" binding:"required"`
}

func (controller *APIController) CreateProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 对用户权限进行验证

	arg := repositories.CreateProductParams{
		ProductName: req.ProductName,
		ProductNum: req.ProductNum,
		ProductImage: req.ProductImage,
		ProductUrl: req.ProductUrl,
	}

	if ok := controller.store.CreateProduct(ctx, arg); !ok {
		err := errors.New("failed to insert product")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// 如果还需要拿到插入后的结果还需要进行一次查询?

	ctx.JSON(http.StatusOK, gin.H{"message": "insert into, ok"})
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (controller *APIController) GetProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := repositories.GetProductParams{
		ID: req.ID,
	}

	var product datamodels.Product
	var err error

	if product, err = controller.store.GetProduct(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}











