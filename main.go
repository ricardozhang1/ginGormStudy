package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func main() {
//	// 获取配置信息
//	config, err := conf.LoadConfig("./conf")
//	if err != nil {
//		log.Fatalf("failed to read config, err: %v", err)
//	}
//
//	// 配置日志信息
//	logger := logging.NewLog(config)
//
//	conn, err := gorm.Open(config.DBDriver, config.DBSource)
//	if err != nil {
//		log.Fatalf("failed to connect to DB, err: %v", err)
//	}
//
//	store := repositories.NewStore(conn, config)
//
//	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
//	if err != nil {
//		log.Fatalf("failed to create token maker, err: %v", err)
//	}
//
//	controller := controllers.NewController(store, tokenMaker, config, logger)
//
//	server, err := services.NewServer(config, store, controller, tokenMaker, logger)
//	if err != nil {
//		log.Fatalf("failed to create services, err: %v", err)
//	}
//
//	if err = server.Start(config.ServerAddress); err != nil {
//		log.Fatal("cannot start server: ", err)
//	}
//}
