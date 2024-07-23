package payment

import (
	"context"

	paymentv1 "github.com/Bookil/Bookil-Proto/gen/golang/payment/v1"
	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment paymentv1.PaymentServiceClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := paymentv1.NewPaymentServiceClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &paymentv1.CreateRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	if err != nil {
		return err
	}
	return nil
}
