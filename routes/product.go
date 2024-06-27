package routes

import (
	"synapsis/handlers"
	"synapsis/pkg/psql"

	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepo := repositories.NewProductRepository(psql.DB)
	productHandler := handlers.NewProductHandler(productRepo)

	e.POST("/product", productHandler.CreateProduct)
	e.GET("/products/category/:category", productHandler.GetProductsByCategory)
}
