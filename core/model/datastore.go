package model

import (
	"errors"
	"fmt"
	"shopy/logger"

	"github.com/google/uuid"
)

type DataStoreIntf interface {
	GetCustomerById(id string) (*Customer, error)
	GetAllCustomers() map[string]Customer
	AddOrder(o *Order) (string, error)
	GetOrderById(id string) (*Order, error)
	GetOrdersByCustomerId(id string) (*[]Order, error)
	UpdateOrder(o *Order) error
	AddProduct(p *Product) (string, error)
	GetAllProducts() map[string]Product
	GetProductById(id string) (*Product, error)
	UpdateProduct(p *Product) error
}

type DataStore struct {
	orders    map[string]Order
	products  map[string]Product
	customers map[string]Customer
}

func NewDataStore() DataStoreIntf {
	return &DataStore{
		orders:    make(map[string]Order),
		products:  getDefaultProducts(),
		customers: getDefaultCustomers(),
	}
}

// Customer service datastore implementations
func (d *DataStore) AddCustomer(c *Customer) (string, error) {
	c.Id = generateUUID()
	if _, ok := d.customers[c.Id]; !ok {
		err := fmt.Errorf("customer id %v already exists", c.Id)
		logger.Error(err)
		return "", err
	}
	d.customers[c.Id] = *c
	return c.Id, nil
}

func (d *DataStore) GetCustomerById(id string) (*Customer, error) {
	customer, ok := d.customers[id]
	if !ok {
		return nil, errors.New("unable to fetch customer details for id: " + id)
	}
	return &customer, nil
}

func (d *DataStore) GetAllCustomers() map[string]Customer {
	return d.customers
}

// Order service datastore implementations
func (d *DataStore) AddOrder(o *Order) (string, error) {
	o.Id = generateUUID()
	if _, ok := d.orders[o.Id]; ok {
		err := fmt.Errorf("order id %v already exists", o.Id)
		logger.Error(err)
		return "", err
	}
	d.orders[o.Id] = *o
	return o.Id, nil
}

func (d *DataStore) GetOrdersByCustomerId(customerId string) (*[]Order, error) {
	ordersList := []Order{}
	if _, ok := d.customers[customerId]; !ok {
		return nil, errors.New("unable to fetch customer details for id: " + customerId)
	}
	for _, order := range d.orders {
		if order.CustomerId == customerId {
			ordersList = append(ordersList, order)
		}
	}
	return &ordersList, nil
}

func (d *DataStore) GetOrderById(id string) (*Order, error) {
	order, ok := d.orders[id]
	if !ok {
		return nil, errors.New("unable to fetch order details for id: " + id)
	}
	return &order, nil
}

func (d *DataStore) UpdateOrder(o *Order) error {
	if _, ok := d.orders[o.Id]; !ok {
		return errors.New("unable to fetch order for id: " + o.Id)
	}
	d.orders[o.Id] = *o
	return nil
}

// Product service datastore implementations
func (d *DataStore) AddProduct(p *Product) (string, error) {
	if _, ok := d.products[p.Id]; ok {
		err := fmt.Errorf("product id %v already exists", p.Id)
		logger.Error(err)
		return "", err
	}
	d.products[p.Id] = *p
	return p.Id, nil
}

func (d *DataStore) GetAllProducts() map[string]Product {
	return d.products
}

func (d *DataStore) GetProductById(id string) (*Product, error) {
	product, ok := d.products[id]
	if !ok {
		return nil, errors.New("unable to fetch product details for id: " + id)
	}
	return &product, nil
}

func (d *DataStore) UpdateProduct(p *Product) error {
	if _, ok := d.products[p.Id]; !ok {
		return errors.New("unable to fetch product details for id: " + p.Id)
	}
	d.products[p.Id] = *p
	return nil
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}
