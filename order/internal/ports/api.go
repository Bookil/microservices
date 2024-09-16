package ports

import "github.com/Bookil/microservices/order/internal/application/core/domain"

type APIPort interface {
	SaveOrder(order *domain.Order) (*domain.Order, error)
	PlaceOrder(orderID domain.OrderID) (*domain.Order, error)
	DeleteOrder(orderID domain.OrderID)error
	UpdateOrder(orderID domain.OrderID ,items []*domain.OrderItem)error
}
