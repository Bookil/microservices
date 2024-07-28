package ports

import "github.com/Bookil/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id domain.OrderID) (*domain.Order, error)
	Save(*domain.Order) (*domain.Order, error)
}
