package mock

import (
	"shopy/core/domain"

	"github.com/stretchr/testify/mock"
)

type OrderSVCMock struct {
	mock.Mock
}

func (c *OrderSVCMock) AddOrder(order *domain.Order) (response *domain.Response) {
	args := c.Called(order)
	argsData := args.Get(0)
	if argsData != nil {
		response = argsData.(*domain.Response)
	}
	return
}

func (c *OrderSVCMock) GetAllOrdersForCustomer(customerId string) (response *domain.Response) {
	args := c.Called(customerId)
	argsData := args.Get(0)
	if argsData != nil {
		response = argsData.(*domain.Response)
	}
	return
}

func (c *OrderSVCMock) UpdateOrderStatus(orderStatus *domain.OrderStatus) (response *domain.Response) {
	args := c.Called(orderStatus)
	argsData := args.Get(0)
	if argsData != nil {
		response = argsData.(*domain.Response)
	}
	return
}
