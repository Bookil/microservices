package api

import (
	"errors"
	"testing"

	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"github.com/Bookil/microservices/order/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Should_Place_Order(t *testing.T) {
	payment := new(mocks.MockedPayment)
	db := new(mocks.MockedDB)

	orderItem := domain.NewOrderItem("item1", "123", 2.5, 1)
	order := domain.NewOrder("1", []*domain.OrderItem{orderItem})

	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(order, nil)

	application := NewApplication(db, payment)

	_, err := application.PlaceOrder(order)

	assert.Nil(t, err)
}
func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := new(mocks.MockedPayment)
	db := new(mocks.MockedDB)
	ErrConnectionFailed := errors.New("connection error")

	orderItem := domain.NewOrderItem("item1", "123", 2.5, 1)
	order := domain.NewOrder("1", []*domain.OrderItem{orderItem})

	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(order, ErrConnectionFailed)

	application := NewApplication(db, payment)

	_, err := application.PlaceOrder(order)

	assert.EqualError(t, err, ErrConnectionFailed.Error())
}

func Test_Should_Return_Error_When_Payment_Fail(t *testing.T) {
	payment := new(mocks.MockedPayment)
	db := new(mocks.MockedDB)
	ErrInsufficientBalance := errors.New("insufficient balance")

	orderItem := domain.NewOrderItem("item1", "123", 2.5, 1)
	order := domain.NewOrder("1", []*domain.OrderItem{orderItem})

	payment.On("Charge", mock.Anything).Return(ErrInsufficientBalance)
	db.On("Save", mock.Anything).Return(order, nil)

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(order)

	st, _ := status.FromError(err)

	assert.Equal(t, st.Message(), "order creation failed")
	assert.Equal(t, st.Code(), codes.InvalidArgument)
}
