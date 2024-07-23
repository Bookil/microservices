package port

import "github.com/Bookil/Bookil-Microservices/payment/internal/application/core/domain"

type APIPort interface {
	Charge(payment *domain.Payment) (*domain.Payment, error)
}
