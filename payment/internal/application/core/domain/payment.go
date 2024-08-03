package domain

import (
	"fmt"
	"time"

	"github.com/Bookil/Bookil-Microservices/payment/utils/random"
)

type (
	PaymentID  = string
	OrderID    = string
	CustomerID = string
	Payment    struct {
		ID         PaymentID  `gorm:"type:varchar(191);not null"`
		CustomerID CustomerID `gorm:"type:varchar(191);not null"`
		OrderID    OrderID    `gorm:"type:varchar(191);not null"`
		TotalPrice float32    `gorm:"=not null"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)

func NewPayment(customerID CustomerID, orderID OrderID, totalPrice float32) *Payment {
	return &Payment{
		ID:         fmt.Sprintf("%d", random.GenerateID()),
		CustomerID: customerID,
		OrderID:    orderID,
		TotalPrice: totalPrice,
	}
}
