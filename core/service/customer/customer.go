package customer

import (
	"net/http"
	"shopy/core/domain"
	"shopy/core/model"
	"shopy/logger"
	"shopy/util"
)

type CustomerSVC struct {
	dataStore model.DataStoreIntf
}

// NewCustomerService initializes and returns the customer svc as per given values
func NewCustomerService(dataStore model.DataStoreIntf) CustomerIntf {
	return &CustomerSVC{
		dataStore: dataStore,
	}
}

// GetAllCustomers gets all the customer details from the datastore
func (p *CustomerSVC) GetAllCustomers() *domain.Response {
	funcName := logger.FuncName()
	logger.Info("inside ", funcName)
	customersMap := p.dataStore.GetAllCustomers()
	customersList := &[]domain.Customer{}
	sortedKeysList := util.SortCustomerIdKeys(customersMap)
	for _, key := range sortedKeysList {
		*customersList = append(*customersList, domain.Customer{
			Id:           customersMap[key].Id,
			Name:         customersMap[key].Name,
			MobileNumber: customersMap[key].MobileNumber,
		})
	}
	return domain.MapResponse(http.StatusOK, util.ConstSuccess, customersList)
}
