package model

import (
	"time"
)

type Order struct {
	Id                   string
	CustomerId           string
	OrderedItems         []OrderedItems
	DeliveryAddress      string
	OrderAmount          float64
	DiscountAmount       float64
	IsDiscountApplicable bool
	Status               string
	OrderedDate          time.Time
	DispatchDate         time.Time
}

type OrderedItems struct {
	ProductId string
	Quantity  int
}
