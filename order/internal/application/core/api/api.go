package api
 
import (
    "github.com/Bookil/microservices/order/internal/application/core/domain"
    "github.com/Bookil/microservices/order/internal/ports"
)
 
type Application struct {
    db ports.DBPort
    payment ports.PaymentPort
}
 
func NewApplication(db ports.DBPort,paymentPort ports.PaymentPort) *Application {
    return &Application{
        db: db,
        payment: paymentPort,
    }
}
 
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
    err := a.db.Save(&order)
    if err != nil {
        return domain.Order{}, err
    }

    err = a.payment.Charge(&order)
    if err != nil {
        return domain.Order{}, err
    } 
    return order, nil
}
