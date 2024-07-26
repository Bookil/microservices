package db

import (
	"fmt"
	"sync"

	"github.com/Bookil/microservices/order/config"
	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

var (
	dbInc *Adapter
	mutex   = new(sync.Mutex)
)
type Adapter struct {
	db *gorm.DB
}

func generateURL(url *config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",url.Username, url.Password, url.Host, url.Port, url.DBName)
}

func NewAdapter(url *config.Mysql) (*Adapter, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if dbInc == nil {
		db, openErr := gorm.Open(mysql.Open(generateURL(url)), &gorm.Config{})
		if openErr != nil {
			return nil, fmt.Errorf("db connection error: %v", openErr)
		}
	
		err := db.AutoMigrate(&Order{}, OrderItem{})
		if err != nil {
			return nil, fmt.Errorf("db migration error: %v", err)
		}
		dbInc = &Adapter{db: db}
	}
	return dbInc, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
