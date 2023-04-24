package order

import (
	"fmt"
	"reflect"
	"shopy/core/domain"
	"shopy/core/model"
	"shopy/util"
	"time"

	"github.com/spf13/viper"
)

// validateProductQuantity validates the ordered quantity as per the requirements and returns error, if any validation fails
func validateProductQuantity(orderedQuantity, availableQuantity int, productId string) error {
	if availableQuantity == 0 {
		return fmt.Errorf("productId %v is not available, please order only available products", productId)
	}
	maxProductsInOrder := viper.GetInt(util.MaxProductsInOrder)
	if orderedQuantity <= 0 || orderedQuantity > maxProductsInOrder {
		return fmt.Errorf("productId %v quantity should be greater than 0 and less than %v", productId, maxProductsInOrder)
	}
	if orderedQuantity > availableQuantity {
		return fmt.Errorf("productId %v is not avaiable in given quantity, please order available quantity or less", productId)
	}
	return nil
}

// convertModelToDomainOrdersList maps the model orders list to domain orders list
func convertModelToDomainOrdersList(ordersList *[]model.Order) *[]domain.Order {
	ordersListResponse := &[]domain.Order{}
	for _, ordersListItem := range *ordersList {
		order := domain.Order{
			Id:              ordersListItem.Id,
			CustomerId:      ordersListItem.CustomerId,
			DeliveryAddress: ordersListItem.DeliveryAddress,
			Status:          ordersListItem.Status,
		}
		order.OrderedDate = ordersListItem.OrderedDate
		if !reflect.DeepEqual(ordersListItem.DispatchDate, time.Time{}) {
			dispatchDate := ordersListItem.DispatchDate
			order.DispatchDate = &dispatchDate
		}
		if ordersListItem.IsDiscountApplicable {
			order.IsDiscountApplied = ordersListItem.IsDiscountApplicable
			order.TotalAmount = util.RoundFloat(ordersListItem.OrderAmount - ordersListItem.DiscountAmount)
		} else {
			order.TotalAmount = util.RoundFloat(ordersListItem.OrderAmount)
		}
		orderedItemsList := []domain.OrderedItems{}
		for _, orderedItemsListItem := range ordersListItem.OrderedItems {
			orderedItemsList = append(orderedItemsList, domain.OrderedItems{
				ProductId: orderedItemsListItem.ProductId,
				Quantity:  orderedItemsListItem.Quantity,
			})
		}
		order.OrderedItems = orderedItemsList
		*ordersListResponse = append(*ordersListResponse, order)
	}
	return ordersListResponse
}
