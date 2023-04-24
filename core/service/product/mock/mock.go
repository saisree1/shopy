package mock

import (
	"shopy/core/domain"

	"github.com/stretchr/testify/mock"
)

type ProductSVCMock struct {
	mock.Mock
}

func (c *ProductSVCMock) AddProduct(product *domain.Product) (response *domain.Response) {
	args := c.Called(product)
	argsData := args.Get(0)
	if argsData != nil {
		response = argsData.(*domain.Response)
	}
	return
}

func (c *ProductSVCMock) GetAllProducts() (response *domain.Response) {
	args := c.Called()
	argsData := args.Get(0)
	if argsData != nil {
		response = argsData.(*domain.Response)
	}
	return
}
