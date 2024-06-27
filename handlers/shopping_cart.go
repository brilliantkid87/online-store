package handlers

import (
	"log"
	"net/http"
	"synapsis/models"
	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

type ShoppingCartHandler struct {
	Repo *repositories.ShoppingCartRepository
}

func NewShopcartHandler(repo *repositories.ShoppingCartRepository) *ShoppingCartHandler {
	return &ShoppingCartHandler{Repo: repo}
}

func (h *ShoppingCartHandler) AddToCart(c echo.Context) error {
	var cart models.ShoppingCart
	if err := c.Bind(&cart); err != nil {
		log.Println("Invalid input:", err)
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	if cart.UserId == "" || cart.ProductId == "" || cart.Quantity <= 0 {
		return c.String(http.StatusBadRequest, "Missing or invalid required parameters")
	}

	err := h.Repo.AddToCart(c.Request().Context(), &cart)
	if err != nil {
		log.Println("Failed to add to cart:", err)
		return c.String(http.StatusInternalServerError, "Failed to add to cart")
	}

	return c.JSON(http.StatusOK, "Item added to cart successfully")
}

func (h *ShoppingCartHandler) GetCartItems(c echo.Context) error {
	userID := c.Param("user_id")

	items, err := h.Repo.GetCartItems(c.Request().Context(), userID)
	if err != nil {
		log.Println("Failed to get cart items:", err)
		return c.String(http.StatusInternalServerError, "Failed to fetch cart items")
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ShoppingCartHandler) CheckoutAndPay(c echo.Context) error {
	var params map[string]interface{}
	if err := c.Bind(&params); err != nil {
		log.Println("Invalid input:", err)
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	transactionID, err := h.Repo.CheckoutAndPay(c.Request().Context(), params)
	if err != nil {
		log.Println("Failed to checkout and pay:", err)
		return c.JSON(http.StatusInternalServerError, "Failed to checkout and pay")
	}

	return c.JSON(http.StatusOK, map[string]string{"transaction_id": transactionID})
}
