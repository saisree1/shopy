package customer

import (
	"net/http"
	"net/http/httptest"
	"shopy/core/domain"
	mockCustomerSVC "shopy/core/service/customer/mock"
	"shopy/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCustomersSuccessScenario(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestContext(w)
	customerSVC := mockCustomerSVC.CustomerSVCMock{}
	customerSVC.On("GetAllCustomers").Return(domain.MapResponse(http.StatusOK, util.ConstSuccess, &[]domain.Customer{
		{Id: "cm1", Name: "Mock", MobileNumber: "9999999999"},
	}))
	GetAllCustomers(&customerSVC)(ctx)
	assert.Equal(t, w.Code, http.StatusOK)
}

func getTestContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/customer", nil)
	c.Request.Header.Add("Content-Type", "application/json")
	return c
}
