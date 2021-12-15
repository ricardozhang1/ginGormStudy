package controllers

import (
	"github.com/gin-gonic/gin"
	"irisStudy/repositories"
)

type Contro interface {
	CreatOrder(ctx *gin.Context)
	CreatUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type APIController struct {
	store repositories.Store
}

func NewController(store repositories.Store) Contro {
	return &APIController{store: store}
}

func errorResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}

