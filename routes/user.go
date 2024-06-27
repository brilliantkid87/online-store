package routes

import (
	"synapsis/handlers"
	"synapsis/pkg/psql"

	"synapsis/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepo := repositories.NewUserRepository(psql.DB)
	userHandler := handlers.NewUserHandler(userRepo)

	e.POST("/register", userHandler.RegisterUser)
	e.POST("/login", userHandler.Login)
}
