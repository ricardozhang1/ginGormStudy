package controllers

import (
	"github.com/gin-gonic/gin"
	"irisStudy/conf"
	"irisStudy/logging"
	"irisStudy/repositories"
	"irisStudy/token"
)

type Contro interface {
	CreatOrder(ctx *gin.Context)
	CreatUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	LogUser(ctx *gin.Context)

	CreateProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)

	GetOrderByProduct(ctx *gin.Context)
	GetOrderByUser(ctx *gin.Context)
}

type APIController struct {
	store repositories.Store
	tokenMaker token.Maker
	config conf.Config
	logger *logging.Log
}

func NewController(store repositories.Store, tokenMaker token.Maker, config conf.Config, logger *logging.Log) Contro {
	return &APIController{store: store, tokenMaker: tokenMaker, config: config, logger: logger}
}

func errorResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}

