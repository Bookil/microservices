package api

import (
	"github.com/Bookil/Bookil-Microservices/payment/internal/application/core/domain"
	"github.com/Bookil/Bookil-Microservices/payment/internal/port"
)


type Application struct {
	db port.DBPort
}

func NewApplication(db port.DBPort)*Application{
    return &Application{
        db: db,
    }
}
func (a *Application) Charge(payment *domain.Payment) (*domain.Payment, error) {
	err := a.db.Save(payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
