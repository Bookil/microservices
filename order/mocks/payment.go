package mocks

import (
	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockedPayment struct {
	mock.Mock
}

func (p *MockedPayment) Charge(order *domain.Order) error {
	args := p.Called(order)
	return args.Error(0)
}

