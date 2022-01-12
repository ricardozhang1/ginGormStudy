package services

import (
	"github.com/gin-gonic/gin"
	"irisStudy/conf"
	"irisStudy/controllers"
	"irisStudy/logging"
	"irisStudy/middleware"
	"irisStudy/repositories"
	"irisStudy/token"
)

type Server struct {
	config conf.Config
	logger *logging.Log
	store repositories.Store
	controller controllers.Contro
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config conf.Config, store repositories.Store, controller controllers.Contro, tokenMaker token.Maker, logger *logging.Log) (*Server, error) {
	server := &Server{
		config: config,
		logger: logger,
		store: store,
		tokenMaker: tokenMaker,
		controller: controller,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// 注册登录API
	router.POST("/user", server.controller.CreatUser)
	router.POST("/login", server.controller.LogUser)

	// 用户操作API
	userRouter := router.Group("/user").Use(middleware.AuthMiddleware(server.tokenMaker))
	{
		userRouter.DELETE("", server.controller.DeleteUser)
	}

	// 订单API
	orderRouter := router.Group("/order").Use(middleware.AuthMiddleware(server.tokenMaker))
	{
		orderRouter.POST("", server.controller.CreatOrder)
		orderRouter.GET("/user/:user_id", server.controller.GetOrderByUser)
		orderRouter.GET("/product/:product_id", server.controller.GetOrderByProduct)
	}

	// 商品API
	productRouter := router.Group("/product").Use(middleware.AuthMiddleware(server.tokenMaker))
	{
		productRouter.POST("", server.controller.CreateProduct)
		productRouter.GET("/:id", server.controller.GetProduct)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}




