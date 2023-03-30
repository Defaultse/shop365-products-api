package v1

import (
	"net/http"

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
		h.GET("/:id", r.productByID)
		h.GET("/category/:id", r.productByCategoryID)
	}
}

func (r *ProductRoutes) allProducts(c *gin.Context) {
	// var request doTranslateRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	// r.l.Error(err, "http - v1 - doTranslate")
	// 	errorResponse(c, http.StatusBadRequest, "invalid request body")

	// 	return
	// }

	r.uc.GetAllProducts()

	c.JSON(http.StatusOK, "")
}

func (r *ProductRoutes) productByID(c *gin.Context) {

}

func (r *ProductRoutes) productByCategoryID(c *gin.Context) {

}
