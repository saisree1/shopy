package order

import (
	"errors"
	"net/http"
	"shopy/core/domain"
	"shopy/core/model"
	modelMock "shopy/core/model/mock"
	"shopy/logger"
	"shopy/util"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.InitLogger()
}

func TestGetAllOrdersForACustomer(t *testing.T) {
	var testCases = []struct {
		scenarioName, customerId string
		responseCode             int
	}{
		{"SuccessScenario", "c1", http.StatusOK},
		{"FailureScenarioForInvalidCustomerId", "c0", http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			orderSVC := getOrderSVCForGetAllOrders(testCaseData.scenarioName, testCaseData.customerId)
			response := orderSVC.GetAllOrdersForCustomer(testCaseData.customerId)
			assert.Equal(t, response.Code, testCaseData.responseCode)
		})
	}
}

func getOrderSVCForGetAllOrders(scenarioName, customerId string) OrderIntf {
	datastoreMock := modelMock.DataStoreMock{}
	if scenarioName == "FailureScenarioForInvalidCustomerId" {
		datastoreMock.On("GetOrdersByCustomerId", mock.Anything).Return(nil, errors.New("invalid customer id"))
	} else {
		datastoreMock.On("GetOrdersByCustomerId", mock.Anything).Return(&[]model.Order{
			{
				Id:                   "m1",
				CustomerId:           customerId,
				IsDiscountApplicable: false,
				OrderAmount:          10.99,
				DeliveryAddress:      "mock Addresss",
				OrderedItems: []model.OrderedItems{
					{
						ProductId: "prod1",
						Quantity:  5,
					},
				},
				Status: util.ConstPlaced,
			},
			{
				Id:                   "m2",
				CustomerId:           customerId,
				IsDiscountApplicable: true,
				OrderAmount:          10.99,
				DiscountAmount:       1.1,
				DeliveryAddress:      "mock Addresss",
				OrderedItems: []model.OrderedItems{
					{
						ProductId: "prod1",
						Quantity:  5,
					},
				},
				DispatchDate: time.Now(),
				Status:       util.ConstDispatched,
			},
		}, nil)
	}
	return NewOrderService(datastoreMock)
}

func TestUpdateOrderStatus(t *testing.T) {
	var testCases = []struct {
		scenarioName, status string
		responseCode         int
	}{
		{"SuccessScenario", util.ConstDispatched, http.StatusOK},
		{"SuccessScenarioForCancelledOrder", util.ConstCancelled, http.StatusOK},
		{"FailureScenarioForInvalidOrderId", util.ConstPlaced, http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			orderSVC := getOrderSVCForUpdateOrderStatus(testCaseData.scenarioName, testCaseData.status)
			response := orderSVC.UpdateOrderStatus(&domain.OrderStatus{
				Id:     "m1",
				Status: testCaseData.status,
			})
			assert.Equal(t, response.Code, testCaseData.responseCode)
		})
	}
}

func getOrderSVCForUpdateOrderStatus(scenarioName, status string) OrderIntf {
	datastoreMock := modelMock.DataStoreMock{}
	if scenarioName == "FailureScenarioForInvalidOrderId" {
		datastoreMock.On("GetOrderById", mock.Anything).Return(nil, errors.New("invalid order id"))
	} else if scenarioName == "SuccessScenarioForCancelledOrder" {
		datastoreMock.On("GetOrderById", mock.Anything).Return(&model.Order{
			Id:         "m1",
			CustomerId: "cm1",
			OrderedItems: []model.OrderedItems{
				{ProductId: "pm1", Quantity: 2},
			},
			Status: util.ConstPlaced,
		}, nil)
		datastoreMock.On("GetProductById", mock.Anything).Return(&model.Product{}, nil)
		datastoreMock.On("UpdateProduct", mock.Anything).Return(nil)
		datastoreMock.On("UpdateOrder", mock.Anything).Return(nil)
	} else {
		datastoreMock.On("GetOrderById", mock.Anything).Return(&model.Order{
			Id:         "m1",
			CustomerId: "cm1",
		}, nil)
		datastoreMock.On("UpdateOrder", mock.Anything).Return(nil)
	}
	return NewOrderService(datastoreMock)
}

func TestAddOrder(t *testing.T) {
	viper.Set(util.MaxProductsInOrder, 10)
	viper.Set(util.DiscountProductsCountConfig, 2)
	viper.Set(util.DiscountProductsCategoryConfig, "PREMIUM")
	defer viper.Reset()
	var testCases = []struct {
		scenarioName string
		responseCode int
	}{
		{"SuccessScenario", http.StatusOK},
		{"SuccessScenarioWithDiscountedItemsInOrder", http.StatusOK},
		{"FailureScenarioForInvalidCustomerId", http.StatusBadRequest},
		{"FailureScenarioForInvalidProductId", http.StatusBadRequest},
		{"FailureScenarioForOrderedProductIsNotAvailable", http.StatusBadRequest},
		{"FailureScenarioForOrderingProductMoreThanMaxLimit", http.StatusBadRequest},
		{"FailureScenarioForOrderingProductMoreThanAvailableQuantity", http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			orderSVC := getOrderSVCForAddOrder(testCaseData.scenarioName)
			order := &domain.Order{
				CustomerId: "cm1",
				OrderedItems: []domain.OrderedItems{
					{ProductId: "pm1", Quantity: 10},
				},
				DeliveryAddress: "mock addresss",
			}
			if testCaseData.scenarioName == "FailureScenarioForOrderingProductMoreThanMaxLimit" {
				order.OrderedItems = []domain.OrderedItems{{ProductId: "pm1", Quantity: 20}}
			} else if testCaseData.scenarioName == "SuccessScenarioWithDiscountedItemsInOrder" {
				order.OrderedItems = append(order.OrderedItems, domain.OrderedItems{ProductId: "pm2", Quantity: 2})
			}
			response := orderSVC.AddOrder(order)
			assert.Equal(t, response.Code, testCaseData.responseCode)
		})
	}
}

func getOrderSVCForAddOrder(scenarioName string) OrderIntf {
	datastoreMock := modelMock.DataStoreMock{}
	if scenarioName == "FailureScenarioForInvalidCustomerId" {
		datastoreMock.On("GetCustomerById", mock.Anything).Return(nil, errors.New("unable to find the customer details"))
	} else {
		datastoreMock.On("GetCustomerById", mock.Anything).Return(&model.Customer{}, nil)
	}
	if scenarioName == "FailureScenarioForInvalidProductId" {
		datastoreMock.On("GetProductById", mock.Anything).Return(nil, errors.New("unable to find the product"))
	} else if scenarioName == "FailureScenarioForOrderedProductIsNotAvailable" {
		datastoreMock.On("GetProductById", mock.Anything).Return(&model.Product{
			Id:       "pm1",
			Price:    10.00,
			Quantity: 0,
			Category: "Regular",
		}, nil)
	} else if scenarioName == "FailureScenarioForOrderingProductMoreThanAvailableQuantity" {
		datastoreMock.On("GetProductById", mock.Anything).Return(&model.Product{
			Id:       "pm1",
			Price:    10.00,
			Quantity: 5,
			Category: "Regular",
		}, nil)
	} else if scenarioName == "SuccessScenarioWithDiscountedItemsInOrder" {
		datastoreMock.On("GetProductById", "pm1").Return(&model.Product{
			Id:       "pm1",
			Price:    10.00,
			Quantity: 20,
			Category: "Premium",
		}, nil).Once()
		datastoreMock.On("GetProductById", "pm2").Return(&model.Product{
			Id:       "pm2",
			Price:    10.00,
			Quantity: 20,
			Category: "Premium",
		}, nil)
	} else {
		datastoreMock.On("GetProductById", mock.Anything).Return(&model.Product{
			Id:       "pm1",
			Price:    10.00,
			Quantity: 20,
			Category: "Regular",
		}, nil)
	}
	datastoreMock.On("UpdateProduct", mock.Anything).Return(nil)
	datastoreMock.On("AddOrder", mock.Anything).Return("om1", nil)
	return NewOrderService(datastoreMock)
}
