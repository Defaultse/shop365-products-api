package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"shop365-products-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type ProductRoutes struct {
	uc usecase.ProductUC
	v  *validator.Validate
}

func NewProductRoutes(handler *gin.RouterGroup, v *validator.Validate, uc usecase.ProductUC) {
	r := &ProductRoutes{
		uc: uc,
		v:  v,
	}

	h := handler.Group("/product")
	{
		h.GET("/", r.allProducts)
		h.GET("/:id/:shard", r.productByID)
	}
}

func (r *ProductRoutes) allProducts(c *gin.Context) {
	// var request doTranslateRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	// r.l.Error(err, "http - v1 - doTranslate")
	// 	errorResponse(c, http.StatusBadRequest, "invalid request body")

	// 	return
	// }

	// r.uc.GetAllProducts()

	c.JSON(http.StatusOK, "")
}

func (r *ProductRoutes) productByID(c *gin.Context) {
	productID := c.Query("id")
	shardID := c.Query("shard")

	fmt.Println(12)

	fmt.Println(productID, shardID)

	productIDint64, err := strconv.ParseInt(productID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	shardIDInt64, err := strconv.ParseInt(shardID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := r.uc.GetByID(productIDint64, shardIDInt64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, product)
}
