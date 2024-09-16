package ports

import "github.com/Bookil/microservices/order/internal/application/core/domain"

type DBPort interface {
	GetOrder(id domain.OrderID) (*domain.Order, error)
	SaveOrder(*domain.Order) (*domain.Order, error)
	UpdateOrder(orderID domain.OrderID,OrderItems *[]domain.OrderItem)(*domain.Order,error)
	DeleteOrder(orderItem domain.OrderID) error
}
