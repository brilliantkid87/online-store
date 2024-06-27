package handlers

import (
	"log"
	"net/http"
	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Repo *repositories.TransactionRepository
}

func NewTransactionHandler(repo *repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{Repo: repo}
}

func (h *TransactionHandler) FindTransactions(c echo.Context) error {
	var params map[string]interface{}
	if err := c.Bind(&params); err != nil {
		log.Println("Invalid input:", err)
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	transactions, err := h.Repo.FindTransactions(c.Request().Context(), params)
	if err != nil {
		log.Println("Failed to find transactions:", err)
		return c.JSON(http.StatusInternalServerError, "Failed to find transactions")
	}

	return c.JSON(http.StatusOK, transactions)
}
