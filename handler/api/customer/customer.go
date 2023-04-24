package customer

import (
	"net/http"
	"shopy/core/service/customer"

	"github.com/gin-gonic/gin"
)

// GetAllCustomer method handles the return of all the customers details through customer service
func GetAllCustomers(customerSVC customer.CustomerIntf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := customerSVC.GetAllCustomers()
		ctx.JSON(http.StatusOK, response)
	}
}
