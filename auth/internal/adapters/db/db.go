package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Adapter struct {
		db *gorm.DB
	}
)

var (
	dbInc *Adapter
	mutex = new(sync.Mutex)
)

func generateURL(url *config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", url.Username, url.Password, url.Host, url.Port, url.DBName)
}

func NewAdapter(url *config.Mysql) (*Adapter, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if dbInc == nil {
		genUrl := generateURL(url)

		log.Println("URL:", genUrl)

		db, openErr := gorm.Open(mysql.Open((genUrl)), &gorm.Config{})
		if openErr != nil {
			return nil, fmt.Errorf("db connection error: %v", openErr)
		}

		err := db.AutoMigrate(&domain.Auth{})
		if err != nil {
			return nil, fmt.Errorf("db migration error: %v", err)
		}
		dbInc = &Adapter{db: db}
	}
	return dbInc, nil
}

func (a *Adapter) Create(auth *domain.Auth)error{
	err := a.db.Create(auth).Error

	if err != nil{
		return err
	}

	return nil
}
