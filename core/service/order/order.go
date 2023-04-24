package order

import (
	"net/http"
	"shopy/core/domain"
	"shopy/core/model"
	"shopy/logger"
	"shopy/util"
	"strings"

	"github.com/spf13/viper"
)

type OrderSVC struct {
	dataStore model.DataStoreIntf
}

// NewOrderService initializes and returns the order svc as per given values
func NewOrderService(dataStore model.DataStoreIntf) OrderIntf {
	return &OrderSVC{
		dataStore: dataStore,
	}
}

// AddOrder adds the order details of a customer and updates product details into datastore
func (o *OrderSVC) AddOrder(order *domain.Order) *domain.Response {
	funcName := logger.FuncName()
	logger.Info("inside ", funcName)
	if _, err := o.dataStore.GetCustomerById(order.CustomerId); err != nil {
		logger.Error(funcName, ":", err.Error())
		return domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid customerId")
	}
	var orderAmount float64
	orderModel := &model.Order{
		CustomerId:      order.CustomerId,
		DeliveryAddress: order.DeliveryAddress,
		OrderedItems:    []model.OrderedItems{},
		Status:          util.ConstPlaced,
		OrderedDate:     util.GetIST(),
	}
	productsCount := viper.GetInt(util.DiscountProductsCountConfig)
	productsCategory := viper.GetString(util.DiscountProductsCategoryConfig)
	discountCategoryProductsCount := 0
	for _, orderedItem := range order.OrderedItems {
		productDetails, err := o.dataStore.GetProductById(orderedItem.ProductId)
		if err != nil {
			logger.Error(funcName, ":", err.Error())
			return domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid productId in orderedItems list")
		}
		if err := validateProductQuantity(orderedItem.Quantity, productDetails.Quantity, orderedItem.ProductId); err != nil {
			logger.Error(funcName, ":", err.Error())
			return domain.MapResponseWithoutBody(http.StatusBadRequest, err.Error())
		}
		if strings.EqualFold(productDetails.Category, productsCategory) {
			discountCategoryProductsCount += 1
		}
		orderAmount += productDetails.Price * float64(orderedItem.Quantity)
		orderModel.OrderedItems = append(orderModel.OrderedItems, model.OrderedItems{
			ProductId: orderedItem.ProductId,
			Quantity:  orderedItem.Quantity,
		})
		productDetails.Quantity = productDetails.Quantity - orderedItem.Quantity
		o.dataStore.UpdateProduct(productDetails)
	}
	orderModel.OrderAmount = util.RoundFloat(orderAmount)
	if discountCategoryProductsCount >= productsCount {
		orderModel.IsDiscountApplicable = true
		discountPercentage := viper.GetInt(util.DiscountConfig)
		discountAmount := (orderAmount * float64(discountPercentage)) / 100
		orderModel.DiscountAmount = util.RoundFloat(discountAmount)
	}
	id, _ := o.dataStore.AddOrder(orderModel)
	return domain.MapResponse(http.StatusOK, util.ConstSuccess, &domain.AddItem{Id: id})
}

// GetAllOrdersForCustomer gets all the order details, for a customer, from datastore
func (o *OrderSVC) GetAllOrdersForCustomer(customerId string) *domain.Response {
	funcName := logger.FuncName()
	logger.Info("inside ", funcName)
	ordersList, err := o.dataStore.GetOrdersByCustomerId(customerId)
	if err != nil {
		logger.Error(funcName, ":", err.Error())
		return domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid customerId")
	}
	ordersListResponse := convertModelToDomainOrdersList(ordersList)
	util.SortOrdersByOrderedDate(*ordersListResponse)
	return domain.MapResponse(http.StatusOK, util.ConstSuccess, &ordersListResponse)
}

// UpdateOrderStatus updates all the order details, for a customer, in datastore
func (o *OrderSVC) UpdateOrderStatus(orderStatus *domain.OrderStatus) *domain.Response {
	funcName := logger.FuncName()
	logger.Info("inside ", funcName)
	order, err := o.dataStore.GetOrderById(orderStatus.Id)
	if err != nil {
		logger.Error(funcName, ":", err.Error())
		return domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid orderId")
	}
	order.Status = orderStatus.Status
	if strings.EqualFold(order.Status, util.ConstDispatched) {
		order.DispatchDate = util.GetIST()
	}
	if strings.EqualFold(order.Status, util.ConstCancelled) {
		for _, orderedItem := range order.OrderedItems {
			productDetails, _ := o.dataStore.GetProductById(orderedItem.ProductId)
			productDetails.Quantity += orderedItem.Quantity
			o.dataStore.UpdateProduct(productDetails)
		}
	}
	o.dataStore.UpdateOrder(order)
	return domain.MapResponseWithoutBody(http.StatusOK, "Success")
}
