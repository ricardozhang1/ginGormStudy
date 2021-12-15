package services

import (
	"github.com/gin-gonic/gin"
	"irisStudy/conf"
	"irisStudy/controllers"
	"irisStudy/repositories"
)

type Server struct {
	config conf.Config
	store repositories.Store
	controller controllers.Contro
	router *gin.Engine
}

func NewServer(config conf.Config, store repositories.Store, controller controllers.Contro) (*Server, error) {
	server := &Server{
		config: config,
		store: store,
		controller: controller,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.POST("", server.controller.CreatUser)
		userRouter.DELETE("", server.controller.DeleteUser)
	}

	orderRouter := router.Group("/order")
	{
		orderRouter.POST("", server.controller.CreatOrder)  // TODO
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}




