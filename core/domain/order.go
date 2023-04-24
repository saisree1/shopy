package domain

import "time"

type Order struct {
	Id                string         `json:"id"`
	CustomerId        string         `json:"customerId" binding:"required,startswith=c"`
	OrderedItems      []OrderedItems `json:"orderedItems" binding:"required,gte=1"`
	DeliveryAddress   string         `json:"deliveryAddress" binding:"required"`
	TotalAmount       float64        `json:"totalAmount,omitempty"`
	IsDiscountApplied bool           `json:"isDiscountApplied,omitempty"`
	Status            string         `json:"status,omitempty"`
	OrderedDate       time.Time      `json:"orderedDate,omitempty"`
	DispatchDate      *time.Time     `json:"dispatchDate,omitempty"`
}

type AddItem struct {
	Id string `json:"id"`
}

type OrderedItems struct {
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type OrderStatus struct {
	Id     string `json:"id" binding:"required"`
	Status string `json:"status" binding:"required,oneof=Placed Dispatched Completed Cancelled"`
}
