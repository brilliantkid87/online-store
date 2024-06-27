package handlers

import (
	"net/http"

	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Repo *repositories.ProductRepository
}

func NewProductHandler(repo *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var params map[string]interface{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	ProductId, err := h.Repo.RegisterProduct(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to register product")
	}

	return c.JSON(http.StatusOK, map[string]string{"product_id": ProductId})
}

func (h *ProductHandler) GetProductsByCategory(c echo.Context) error {
	category := c.Param("category")

	products, err := h.Repo.GetProductsByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}
