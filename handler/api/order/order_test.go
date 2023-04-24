package order

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"shopy/core/domain"
	mockOrderSVC "shopy/core/service/order/mock"
	"shopy/logger"
	"shopy/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.InitLogger()
}

func TestGetAllOrdersForCustomer(t *testing.T) {
	var testCases = []struct {
		scenarioName, customerId string
		responseCode             int
	}{
		{"SuccessScenario", "cm1", http.StatusOK},
		{"FailureScenarioForInvalidCustomerId", "", http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := getTestContextForGetAllOrdersForCustomer(w, testCaseData.customerId)
			orderSVC := mockOrderSVC.OrderSVCMock{}
			orderSVC.On("GetAllOrdersForCustomer", mock.Anything).Return(domain.MapResponse(http.StatusOK, util.ConstSuccess, &[]domain.Order{
				{Id: "cm1", CustomerId: "cm1", OrderedItems: []domain.OrderedItems{{ProductId: "pm1", Quantity: 1}}, DeliveryAddress: "mock address", TotalAmount: 200.92, OrderedDate: util.GetIST(), Status: util.ConstPlaced},
			}))
			GetAllOrdersForCustomer(&orderSVC)(ctx)
			assert.Equal(t, w.Code, http.StatusOK)
		})
	}
}

func getTestContextForGetAllOrdersForCustomer(w *httptest.ResponseRecorder, customerId string) *gin.Context {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/order", nil)
	queryParams := url.Values{}
	queryParams.Add(util.ConstCustomerId, customerId)
	c.Request.URL.RawQuery = queryParams.Encode()
	c.Request.Header.Add("Content-Type", "application/json")
	return c
}

func TestAddOrder(t *testing.T) {
	var testCases = []struct {
		scenarioName string
		responseCode int
	}{
		{"SuccessScenario", http.StatusOK},
		{"FailureScenarioForInvalidRequestBody", http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := getTestContextForAddOrder(w, testCaseData.scenarioName)
			orderSVC := mockOrderSVC.OrderSVCMock{}
			orderSVC.On("AddOrder", mock.Anything).Return(domain.MapResponse(http.StatusOK, util.ConstSuccess, &domain.Order{Id: "om1"}))
			AddOrder(&orderSVC)(ctx)
			assert.Equal(t, w.Code, testCaseData.responseCode)
		})
	}
}

func getTestContextForAddOrder(w *httptest.ResponseRecorder, scenarioName string) *gin.Context {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	var orderRequestBytes []byte
	if scenarioName == "FailureScenarioForInvalidRequestBody" {
		orderRequestBytes = []byte(``)
	} else {
		orderRequest := &domain.Order{CustomerId: "cm1", OrderedItems: []domain.OrderedItems{{ProductId: "pm1", Quantity: 2}}, DeliveryAddress: "mockAddress"}
		orderRequestBytes, _ = json.Marshal(orderRequest)
	}
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/order", io.NopCloser(bytes.NewBuffer([]byte(orderRequestBytes))))
	c.Request.Header.Add("Content-Type", "application/json")
	return c
}

func TestUpdateOrderStatus(t *testing.T) {
	var testCases = []struct {
		scenarioName string
		responseCode int
	}{
		{"SuccessScenario", http.StatusOK},
		{"FailureScenarioForInvalidRequestBody", http.StatusBadRequest},
	}
	for _, testCaseData := range testCases {
		t.Run(testCaseData.scenarioName, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := getTestContextForUpdateOrderStatus(w, testCaseData.scenarioName)
			orderSVC := mockOrderSVC.OrderSVCMock{}
			orderSVC.On("UpdateOrderStatus", mock.Anything).Return(domain.MapResponse(http.StatusOK, util.ConstSuccess, &domain.Order{Id: "om1"}))
			UpdateOrderStatus(&orderSVC)(ctx)
			assert.Equal(t, w.Code, testCaseData.responseCode)
		})
	}
}

func getTestContextForUpdateOrderStatus(w *httptest.ResponseRecorder, scenarioName string) *gin.Context {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	var orderRequestBytes []byte
	if scenarioName == "FailureScenarioForInvalidRequestBody" {
		orderRequestBytes = []byte(``)
	} else {
		orderRequest := &domain.OrderStatus{Id: "om1", Status: util.ConstDispatched}
		orderRequestBytes, _ = json.Marshal(orderRequest)
	}
	c.Request, _ = http.NewRequest(http.MethodPut, "/api/v1/order/status", io.NopCloser(bytes.NewBuffer([]byte(orderRequestBytes))))
	c.Request.Header.Add("Content-Type", "application/json")
	return c
}
