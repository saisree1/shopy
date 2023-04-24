package product

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"shopy/core/domain"
	mockProductSVC "shopy/core/service/product/mock"
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

func TestAddProduct(t *testing.T) {
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
			ctx := getTestContextForAddProduct(w, testCaseData.scenarioName)
			productSVC := mockProductSVC.ProductSVCMock{}
			productSVC.On("AddProduct", mock.Anything).Return(domain.MapResponse(http.StatusOK, util.ConstSuccess, &domain.Order{Id: "om1"}))
			AddProduct(&productSVC)(ctx)
			assert.Equal(t, w.Code, testCaseData.responseCode)
		})
	}
}

func getTestContextForAddProduct(w *httptest.ResponseRecorder, scenarioName string) *gin.Context {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	var orderRequestBytes []byte
	if scenarioName == "FailureScenarioForInvalidRequestBody" {
		orderRequestBytes = []byte(``)
	} else {
		orderRequest := &domain.Product{Name: "mock Product", Price: 20.00, Category: "Regular", Quantity: 10}
		orderRequestBytes, _ = json.Marshal(orderRequest)
	}
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/order", io.NopCloser(bytes.NewBuffer([]byte(orderRequestBytes))))
	c.Request.Header.Add("Content-Type", "application/json")
	return c
}

func TestGetAllProductsSuccessScenario(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestContext(w)
	productSVC := mockProductSVC.ProductSVCMock{}
	productSVC.On("GetAllProducts").Return(domain.MapResponse(http.StatusOK, util.ConstSuccess, &[]domain.Product{
		{Id: "cm1", Name: "Mock", Category: "Regular"},
	}))
	GetAllProducts(&productSVC)(ctx)
	assert.Equal(t, w.Code, http.StatusOK)
}

func getTestContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/product", nil)
	c.Request.Header.Add("Content-Type", "application/json")
	return c
}
