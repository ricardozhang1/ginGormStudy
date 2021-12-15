package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"irisStudy/conf"
	"irisStudy/controllers"
	"irisStudy/repositories"
	"irisStudy/services"
	"log"
)

func main() {
	// 获取配置信息
	config, err := conf.LoadConfig("./conf")
	if err != nil {
		log.Fatalf("failed to read config, err: %v", err)
	}

	conn, err := gorm.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("failed to connect to DB, err: %v", err)
	}

	store := repositories.NewStore(conn)
	controller := controllers.NewController(store)

	server, err := services.NewServer(config, store, controller)
	if err != nil {
		log.Fatalf("failed to create services, err: %v", err)
	}

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
