package mock

import (
	"shopy/core/domain"

	"github.com/stretchr/testify/mock"
)

type CustomerSVCMock struct {
	mock.Mock
}

func (c *CustomerSVCMock) GetAllCustomers() (response *domain.Response) {
	args := c.Called()
	argsData := args.Get(0)
	if argsData != nil {
		response = argsData.(*domain.Response)
	}
	return
}
