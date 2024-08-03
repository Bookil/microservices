package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/Bookil/microservices/order/config"
	"github.com/Bookil/microservices/order/internal/application/core/domain"
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

		err := db.AutoMigrate(&domain.Order{}, &domain.OrderItem{})
		if err != nil {
			return nil, fmt.Errorf("db migration error: %v", err)
		}
		dbInc = &Adapter{db: db}
	}
	return dbInc, nil
}

func (a Adapter) Get(id string) (*domain.Order, error) {
	order := &domain.Order{}
	res := a.db.First(order, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) (*domain.Order, error) {
	res := a.db.Create(order)
	if res.Error == nil {
		return nil, res.Error
	}

	return order, nil
}
