package product

import (
	"errors"
	"net/http"
	"shopy/core/domain"
	"shopy/core/model"
	modelMock "shopy/core/model/mock"
	"shopy/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.InitLogger()
}

func TestAddOrder(t *testing.T) {
	var testCases = []struct {
		scenarioName string
		responseCode int
	}{
		{"SuccessScenario", http.StatusOK},
		{"FailureScenarioForAddOrderIntoDatastore", http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			productSVC := getProductSVCForAddProduct(testCaseData.scenarioName)
			response := productSVC.AddProduct(&domain.Product{
				Name:     "Mock product",
				Price:    20.00,
				Category: "Regular",
				Quantity: 20,
			})
			assert.Equal(t, testCaseData.responseCode, response.Code)
		})
	}
}

func getProductSVCForAddProduct(scenarioName string) ProductIntf {
	datastoreMock := modelMock.DataStoreMock{}
	if scenarioName == "FailureScenarioForAddOrderIntoDatastore" {
		datastoreMock.On("AddProduct", mock.Anything).Return("", errors.New("product id already exists"))
	} else {
		datastoreMock.On("AddProduct", mock.Anything).Return("pm1", nil)
	}
	return NewProductService(datastoreMock)
}

func TestGetAllProductsSuccessScenario(t *testing.T) {
	datastoreMock := modelMock.DataStoreMock{}
	datastoreMock.On("GetAllProducts").Return(map[string]model.Product{
		"m1": {
			Id:   "m1",
			Name: "mock1",
		},
	})
	productSVC := NewProductService(datastoreMock)
	response := productSVC.GetAllProducts()
	assert.Equal(t, response.Code, http.StatusOK)
}
