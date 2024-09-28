package db

import (
	"fmt"
	"log"
	"sync"

	"product/config"
	"product/internal/application/core/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
	mutex  = new(sync.Mutex)
)

func generateURL(config *config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DBName)
}

func NewDB(config *config.Mysql) (*gorm.DB, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if gormDB == nil {
		genUrl := generateURL(config)

		log.Println("URL:", genUrl)

		db, openErr := gorm.Open(mysql.Open((genUrl)), &gorm.Config{})
		if openErr != nil {
			return nil, fmt.Errorf("db connection error: %v", openErr)
		}

		err := db.AutoMigrate(&domain.Book{}, &domain.Genre{}, &domain.Author{})
		if err != nil {
			return nil, fmt.Errorf("db migration error: %v", err)
		}

		gormDB = db
	}
	return gormDB, nil
}
