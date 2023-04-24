package mock

import (
	"shopy/core/model"

	"github.com/stretchr/testify/mock"
)

type DataStoreMock struct {
	mock.Mock
}

// GetAllCustomers used for unit test cases for mocking
func (dsMock DataStoreMock) GetAllCustomers() (customersMap map[string]model.Customer) {
	args := dsMock.Called()
	if argsData := args.Get(0); argsData != nil {
		customersMap = argsData.(map[string]model.Customer)
		return
	}
	return
}

// GetCustomerById used for unit test cases for mocking
func (dsMock DataStoreMock) GetCustomerById(id string) (customer *model.Customer, err error) {
	args := dsMock.Called(id)
	if argsData := args.Get(0); argsData != nil {
		customer = argsData.(*model.Customer)
	}
	return customer, args.Error(1)
}

// AddOrder used for unit test cases for mocking
func (dsMock DataStoreMock) AddOrder(o *model.Order) (string, error) {
	args := dsMock.Called(o)
	return args.String(0), args.Error(1)
}

// GetOrderById used for unit test cases for mocking
func (dsMock DataStoreMock) GetOrderById(id string) (order *model.Order, err error) {
	args := dsMock.Called(id)
	if argsData := args.Get(0); argsData != nil {
		order = argsData.(*model.Order)
	}
	return order, args.Error(1)
}

// GetOrdersByCustomerId used for unit test cases for mocking
func (dsMock DataStoreMock) GetOrdersByCustomerId(id string) (ordersList *[]model.Order, err error) {
	args := dsMock.Called(id)
	if argsData := args.Get(0); argsData != nil {
		ordersList = argsData.(*[]model.Order)
	}
	return ordersList, args.Error(1)
}

// UpdateOrder used for unit test cases for mocking
func (dsMock DataStoreMock) UpdateOrder(o *model.Order) error {
	args := dsMock.Called(o)
	return args.Error(0)
}

func (dsMock DataStoreMock) AddProduct(p *model.Product) (string, error) {
	args := dsMock.Called(p)
	return args.String(0), args.Error(1)
}

// GetProductById used for unit test cases for mocking
func (dsMock DataStoreMock) GetProductById(id string) (product *model.Product, err error) {
	args := dsMock.Called(id)
	if argsData := args.Get(0); argsData != nil {
		product = argsData.(*model.Product)
	}
	return product, args.Error(1)
}

// GetAllProducts used for unit test cases for mocking
func (dsMock DataStoreMock) GetAllProducts() (productsMap map[string]model.Product) {
	args := dsMock.Called()
	if argsData := args.Get(0); argsData != nil {
		productsMap = argsData.(map[string]model.Product)
		return
	}
	return
}

// UpdateProduct used for unit test cases for mocking
func (dsMock DataStoreMock) UpdateProduct(p *model.Product) error {
	args := dsMock.Called(p)
	return args.Error(0)
}
