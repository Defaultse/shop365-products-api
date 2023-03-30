package v1

import (
	"shop365-products-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type CategoryRoutes struct {
	uc usecase.CategoryUC
	v  *validator.Validate
}

func NewCategoryRoutes(handler *gin.RouterGroup, v *validator.Validate, uc usecase.CategoryUC) {
	r := &CategoryRoutes{
		uc: uc,
		v:  v,
	}

	h := handler.Group("/category")
	{
		h.GET("/", r.allCategories)
		h.GET("/:id", r.nestedCategoriesByID)
	}
}

func (cr *CategoryRoutes) allCategories(c *gin.Context) {
	categories, err := cr.uc.GetAll()

	if err != nil {
		return
	}

	c.JSON(200, categories)
}

func (cr *CategoryRoutes) nestedCategoriesByID(c *gin.Context) {

}
