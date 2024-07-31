package grpc

import (
	"context"

	orderv1 "github.com/Bookil/Bookil-Proto/gen/golang/order/v1"
	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a Adapter) Create(ctx context.Context, request *orderv1.CreateRequest) (*orderv1.CreateResponse, error) {
	orderItems := []*domain.OrderItem{}
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, &domain.OrderItem{
			ID:          orderItem.ItemId,
			Name:        orderItem.Name,
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	newOrder := domain.NewOrder(request.CustomerId, orderItems)

	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, ErrFailedPlaceOrder
	}

	orderItemResponse := []*orderv1.Item{}
	for _, orderItem := range result.OrderItems {
		orderItemResponse = append(orderItemResponse, &orderv1.Item{
			ItemId:      orderItem.ID,
			Name:        orderItem.Name,
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	orderResponse := &orderv1.Order{
		OrderId:    result.ID,
		CustomerId: result.CustomerID,
		CreatedAt:  timestamppb.New(result.CreatedAt),
		UpdatedAt:  timestamppb.New(result.UpdatedAt),
		TotalPrice: result.TotalPrice(),
		Status:     result.Status,
		Items:      orderItemResponse,
	}
	return &orderv1.CreateResponse{Order: orderResponse}, nil
}
