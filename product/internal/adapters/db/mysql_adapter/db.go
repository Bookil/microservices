package mysql_adapter

import (
	"gorm.io/gorm"
)

type (
	Adapter struct {
		db *gorm.DB
	}
)

func NewAdapter(db *gorm.DB) *Adapter {
	return &Adapter{
		db: db,
	}
}

