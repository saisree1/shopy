package order

import (
	"net/http"
	"shopy/core/domain"
	"shopy/core/service/order"
	"shopy/logger"
	"shopy/util"

	"github.com/gin-gonic/gin"
)

// AddOrder method handles addition of order details through order service
func AddOrder(orderSVC order.OrderIntf) gin.HandlerFunc {
	funcName := logger.FuncName()
	return func(ctx *gin.Context) {
		order := &domain.Order{}
		if err := ctx.BindJSON(order); err != nil {
			logger.Error(funcName, " : invalid request body :: ", err.Error())
			response := domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid RequestBody")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := orderSVC.AddOrder(order)
		ctx.JSON(http.StatusOK, response)
	}
}

// GetAllOrdersForCustomer method handles return of all order details for a customer through order service
func GetAllOrdersForCustomer(orderSVC order.OrderIntf) gin.HandlerFunc {
	funcName := logger.FuncName()
	return func(ctx *gin.Context) {
		customerId := ctx.Query(util.ConstCustomerId)
		if customerId == "" {
			logger.Error(funcName, " : customerId should not be empty")
			response := domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid CustomerId")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := orderSVC.GetAllOrdersForCustomer(customerId)
		ctx.JSON(http.StatusOK, response)
	}
}

// UpdateOrderStatus method updates the status of a order through order service
func UpdateOrderStatus(orderSVC order.OrderIntf) gin.HandlerFunc {
	funcName := logger.FuncName()
	return func(ctx *gin.Context) {
		orderStatus := &domain.OrderStatus{}
		if err := ctx.BindJSON(orderStatus); err != nil {
			logger.Error(funcName, " : invalid request body :: ", err.Error())
			response := domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid RequestBody")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := orderSVC.UpdateOrderStatus(orderStatus)
		ctx.JSON(http.StatusOK, response)
	}
}
