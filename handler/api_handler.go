package handler

import (
	"shopy/core/model"
	"shopy/core/service/customer"
	"shopy/core/service/order"
	"shopy/core/service/product"
	customerHandler "shopy/handler/api/customer"
	orderHandler "shopy/handler/api/order"
	productHandler "shopy/handler/api/product"

	"shopy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAppRoutes(engine *gin.Engine) {
	// model initialization
	dataStore := model.NewDataStore()

	// middleware for making responses to be returned as json type only
	engine.Use(middleware.JSONMiddleware())

	// product svc initialization
	productSVC := product.NewProductService(dataStore)
	addProduct := productHandler.AddProduct(productSVC)
	getAllProducts := productHandler.GetAllProducts(productSVC)
	// customer svc initialization
	customerSVC := customer.NewCustomerService(dataStore)
	getAllCustomers := customerHandler.GetAllCustomers(customerSVC)
	// order svc initialization
	orderSVC := order.NewOrderService(dataStore)
	addOrder := orderHandler.AddOrder(orderSVC)
	getOrders := orderHandler.GetAllOrdersForCustomer(orderSVC)
	updateOrderStatus := orderHandler.UpdateOrderStatus(orderSVC)

	// creating apis and assigning to handlers
	routerGroup := engine.Group("/api/v1")
	{
		// product related apis
		routerGroup.POST("/product", addProduct)
		routerGroup.GET("/product", getAllProducts)

		// customer related apis
		routerGroup.GET("/customer", getAllCustomers)

		// order related apis
		routerGroup.POST("/order", addOrder)
		routerGroup.GET("/order", getOrders)
		routerGroup.PUT("/order/status", updateOrderStatus)
	}
}
