package customer

import (
	"net/http"
	"shopy/core/model"
	"shopy/core/model/mock"
	"shopy/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	logger.InitLogger()
}

func TestGetAllCustomersSuccessScenario(t *testing.T) {
	datastoreMock := mock.DataStoreMock{}
	datastoreMock.On("GetAllCustomers").Return(map[string]model.Customer{
		"m1": {
			Id:   "m1",
			Name: "mock1",
		},
	})
	customerSVC := NewCustomerService(datastoreMock)
	response := customerSVC.GetAllCustomers()
	assert.Equal(t, response.Code, http.StatusOK)
}
