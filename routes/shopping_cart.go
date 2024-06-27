package routes

import (
	"synapsis/handlers"
	"synapsis/pkg/middleware"
	"synapsis/pkg/psql"

	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

func ShoppingCartRoutes(e *echo.Group) {
	cartRepo := repositories.NewShoppingCartRepository(psql.DB)
	cartHandler := handlers.NewShopcartHandler(cartRepo)

	e.POST("/add-to-cart", middleware.Auth(cartHandler.AddToCart))
	e.GET("/get-cart-user/:user_id", middleware.Auth(cartHandler.GetCartItems))
	e.POST("/checkout", middleware.Auth(cartHandler.CheckoutAndPay))
	e.POST("/delete-product", middleware.Auth(cartHandler.RemoveProductInCart))
}
