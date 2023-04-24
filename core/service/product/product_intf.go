package product

import "shopy/core/domain"

type ProductIntf interface {
	AddProduct(product *domain.Product) *domain.Response
	GetAllProducts() *domain.Response
}
