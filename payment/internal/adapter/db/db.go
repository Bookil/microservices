package db

import (
	"fmt"

	"github.com/Bookil/Bookil-Microservices/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserId int32
	OrdrId int32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) Save(payment *domain.Payment) error {
	paymentModel := &Payment{
		UserId: payment.UserId,
		OrdrId: payment.OrderId,
	}

	result := a.db.Create(&paymentModel)
	if result.Error == nil {
		payment.ID = int64(paymentModel.ID)
	}
	return result.Error
}
