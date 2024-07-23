package port

import "github.com/Bookil/Bookil-Microservices/payment/internal/application/core/domain"

type DBPort interface {
	Save(*domain.Payment) error
}