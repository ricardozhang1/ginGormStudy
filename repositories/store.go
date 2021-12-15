package repositories

import (
	"github.com/jinzhu/gorm"
)


type Store interface {
	Querier
}

type SQLStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return &SQLStore{
		db: db,
	}
}

