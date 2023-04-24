package product

import (
	"net/http"
	"shopy/core/domain"
	"shopy/core/model"
	"shopy/logger"
	"shopy/util"
)

type ProductSVC struct {
	dataStore model.DataStoreIntf
}

// NewProductService initializes and returns the product svc as per given values
func NewProductService(dataStore model.DataStoreIntf) ProductIntf {
	return &ProductSVC{
		dataStore: dataStore,
	}
}

// AddProduct adds the new product details into the datastore
func (p *ProductSVC) AddProduct(product *domain.Product) *domain.Response {
	funcName := logger.FuncName()
	logger.Info("inside ", funcName)
	id, err := p.dataStore.AddProduct(&model.Product{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
		Quantity: product.Quantity,
	})
	if err != nil {
		logger.Error(funcName, ":", err.Error())
		return domain.MapResponseWithoutBody(http.StatusBadRequest, "Id already exists")
	}
	return domain.MapResponse(http.StatusOK, util.ConstSuccess, &domain.AddItem{Id: id})
}

// GetAllProducts gets all the products from the datastore
func (p *ProductSVC) GetAllProducts() *domain.Response {
	funcName := logger.FuncName()
	logger.Info("inside ", funcName)
	productsMap := p.dataStore.GetAllProducts()
	productsList := &[]domain.Product{}
	sortedKeysList := util.SortProductIdKeys(productsMap)
	for _, key := range sortedKeysList {
		*productsList = append(*productsList, domain.Product{
			Id:       productsMap[key].Id,
			Name:     productsMap[key].Name,
			Price:    productsMap[key].Price,
			Category: productsMap[key].Category,
			Quantity: productsMap[key].Quantity,
		})
	}
	return domain.MapResponse(http.StatusOK, util.ConstSuccess, productsList)
}
