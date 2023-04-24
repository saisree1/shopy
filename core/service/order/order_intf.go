package order

import "shopy/core/domain"

type OrderIntf interface {
	AddOrder(order *domain.Order) *domain.Response
	GetAllOrdersForCustomer(customerId string) *domain.Response
	UpdateOrderStatus(orderStatus *domain.OrderStatus) *domain.Response
}
