package routes

import (
	"synapsis/handlers"
	"synapsis/pkg/middleware"
	"synapsis/pkg/psql"

	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepo := repositories.NewTransactionRepository(psql.DB)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)

	e.GET("/transactions", middleware.Auth(transactionHandler.FindTransactions))
}
