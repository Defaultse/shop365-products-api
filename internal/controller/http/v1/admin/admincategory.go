package admin

import "github.com/gin-gonic/gin"

type AdminCategoryRoutes struct {
}

func NewAdminCategoryRoutes(handler *gin.RouterGroup) {
	r := &AdminCategoryRoutes{}

	h := handler.Group("/admin/category/")
	{
		h.POST("/", r.createNewCategory)
	}
}

func (r *AdminCategoryRoutes) createNewCategory(c *gin.Context) {

}
