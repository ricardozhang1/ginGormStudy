package repositories

import (
	"github.com/jinzhu/gorm"
	"irisStudy/conf"
	"irisStudy/datamodels"
	"time"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB, config conf.Config) Store {
	// 对Mysql进行配置, 数据迁移
	//db.SingularTable(config.SingularTable)
	db.AutoMigrate(&datamodels.User{})
	db.AutoMigrate(&datamodels.Product{})
	db.AutoMigrate(&datamodels.Order{})

	db.LogMode(config.LogMode) //打印sql语句
	//开启连接池
	db.DB().SetMaxIdleConns(config.MaxIdleConns)                      //最大空闲连接
	db.DB().SetMaxOpenConns(config.MaxOpenConns)                      //最大连接数
	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime)) //最大生存时间(s)

	return &SQLStore{
		db: db,
	}
}
