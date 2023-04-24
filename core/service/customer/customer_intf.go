package customer

import "shopy/core/domain"

type CustomerIntf interface {
	GetAllCustomers() *domain.Response
}
