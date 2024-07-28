package ports

import "github.com/Bookil/microservices/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order *domain.Order) (*domain.Order, error)
}
