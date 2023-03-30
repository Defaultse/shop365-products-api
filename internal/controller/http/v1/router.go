// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/go-playground/validator.v9"

	// Swagger docs.
	_ "shop365-products-api/docs"
	"shop365-products-api/internal/controller/http/v1/admin"
	"shop365-products-api/internal/usecase"
	"shop365-products-api/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, v *validator.Validate, t usecase.AllUseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		NewCategoryRoutes(h, v, t.CategoryUC)
		NewProductRoutes(h, v, t.ProductUC)
		admin.NewAdminCategoryRoutes(h)
		admin.NewAdminProductRoutes(h, v, t.AdminProductUC)
	}
}
