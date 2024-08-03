package grpc

import (
	"context"

	"github.com/Bookil/Bookil-Microservices/payment/internal/application/core/domain"
	paymentv1 "github.com/Bookil/Bookil-Proto/gen/golang/payment/v1"
)

func (a Adapter) Create(ctx context.Context, request *paymentv1.CreateRequest) (*paymentv1.CreateResponse, error) {
	newPayment := domain.NewPayment(request.CustomerId, request.OrderId, request.TotalPrice)

	_, err := a.api.Charge(newPayment)
	if err != nil {
		return nil, ErrFailedChargeCreditCard
	}
	return &paymentv1.CreateResponse{}, nil
}
