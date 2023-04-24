package product

import (
	"net/http"
	"shopy/core/domain"
	"shopy/core/service/product"
	"shopy/logger"

	"github.com/gin-gonic/gin"
)

// AddProduct method handles addition of new product through product service
func AddProduct(productSVC product.ProductIntf) gin.HandlerFunc {
	funcName := logger.FuncName()
	return func(ctx *gin.Context) {
		product := &domain.Product{}
		if err := ctx.BindJSON(product); err != nil {
			logger.Error(funcName, " : invalid request body :: ", err.Error())
			response := domain.MapResponseWithoutBody(http.StatusBadRequest, "Invalid RequestBody")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := productSVC.AddProduct(product)
		ctx.JSON(http.StatusOK, response)
	}
}

// GetAllProducts method returns all the product details through product service
func GetAllProducts(productSVC product.ProductIntf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := productSVC.GetAllProducts()
		ctx.JSON(http.StatusOK, response)
	}
}
