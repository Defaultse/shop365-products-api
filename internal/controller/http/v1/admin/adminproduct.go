package admin

import (
	"net/http"
	"shop365-products-api/internal/dto/admindto"
	"shop365-products-api/internal/usecase/adminuc"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type AdminProductRoutes struct {
	uc adminuc.AdminProductUC
	v  *validator.Validate
}

func NewAdminProductRoutes(handler *gin.RouterGroup, v *validator.Validate, uc adminuc.AdminProductUC) {
	r := &AdminProductRoutes{
		uc: uc,
		v:  v,
	}

	h := handler.Group("/admin/product/")
	{
		h.POST("/", r.createNewProduct)
	}
}

func (r *AdminProductRoutes) createNewProduct(c *gin.Context) {
	product := &admindto.Product{}

	if err := c.ShouldBind(product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.uc.CreateNew(product); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, "cock")
}
