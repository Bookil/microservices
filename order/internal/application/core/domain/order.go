package domain

import (
	"fmt"
	"time"

	"github.com/Bookil/microservices/order/utils/random"
)

type Status = uint32

const (
	Pending Status = iota
	Processing
	Shipped
)

type (
	OrderID    = string
	CustomerID = string
	ItemID     = string
	OrderItem  struct {
		ID          ItemID  `json:"id"`
		Name        string  `json:"name"`
		ProductCode string  `json:"product_code"`
		UnitPrice   float32 `json:"unit_price"`
		Quantity    uint32  `json:"quantity"`
		CreatedAt   int64   `json:"created_at"`
		UpdatedAt   int64   `json:"updated_at"`
	}
)

type Order struct {
	ID         OrderID      `json:"id"`
	CustomerID CustomerID   `json:"customer_id"`
	Status     Status       `json:"status"`
	OrderItems []*OrderItem `json:"order_items"`
	CreatedAt  int64        `json:"created_at"`
	UpdatedAt  int64        `json:"updated_at"`
}

func NewOrder(customerId CustomerID, orderItems []*OrderItem) *Order {
	return &Order{
		ID:         fmt.Sprintf("%d", random.GenerateUserID()),
		CreatedAt:  time.Now().Unix(),
		Status:     Pending,
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}

func (order *Order) BeforeUpdate() {
	order.UpdatedAt = time.Now().Unix()
}

func (order *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, item := range order.OrderItems {
		totalPrice += item.UnitPrice * float32(item.Quantity)
	}
	return totalPrice
}

func NewOrderItem(name, productCode string, unitPrice float32, quantity uint32) *OrderItem {
	return &OrderItem{
		ID:          fmt.Sprintf("%d", random.GenerateUserID()),
		Name:        name,
		CreatedAt:   time.Now().Unix(),
		ProductCode: productCode,
		UnitPrice:   unitPrice,
		Quantity:    quantity,
	}
}

func (item *OrderItem) BeforeUpdate() {
	item.UpdatedAt = time.Now().Unix()
}
