package util

import (
	"shopy/core/domain"
	"shopy/core/model"
	"sort"
	"strconv"
	"strings"
)

func SortProductIdKeys(productsMap map[string]model.Product) []string {
	keys := make([]string, 0)
	for k := range productsMap {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		iVal, _ := strconv.Atoi(strings.Split(keys[i], "p")[1])
		jVal, _ := strconv.Atoi(strings.Split(keys[j], "p")[1])
		return iVal < jVal
	})
	return keys
}

func SortCustomerIdKeys(customersMap map[string]model.Customer) []string {
	keys := make([]string, 0)
	for k := range customersMap {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		iVal, _ := strconv.Atoi(strings.Split(keys[i], "c")[1])
		jVal, _ := strconv.Atoi(strings.Split(keys[j], "c")[1])
		return iVal < jVal
	})
	return keys
}

func SortOrdersByOrderedDate(ordersList []domain.Order) {
	sort.SliceStable(ordersList, func(i, j int) bool {
		return ordersList[i].OrderedDate.Before(ordersList[j].OrderedDate)
	})
}
