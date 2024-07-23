package grpc

import (
	"context"

	orderv1 "github.com/Bookil/Bookil-Proto/gen/golang/order/v1"
	"github.com/Bookil/microservices/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *orderv1.CreateRequest) (*orderv1.CreateResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    int32(orderItem.Quantity),
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)

	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, ErrFailedPlaceOrder
	}
	return &orderv1.CreateResponse{OrderId: result.ID}, nil
}
