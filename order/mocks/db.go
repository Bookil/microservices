package mocks

import (
	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (d *MockedDB) Save(order *domain.Order) (*domain.Order,error){
	args := d.Called(order)
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (d *MockedDB) Get(id string) (*domain.Order, error) {
	args := d.Called(id)
	return args.Get(0).(*domain.Order), args.Error(1)
}
