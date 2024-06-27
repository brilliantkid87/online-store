package main

import (
	"synapsis/database"
	"synapsis/pkg/psql"
	"synapsis/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Connect to the database
	psql.DatabaseConnection()
	database.RunMigration()

	// Group routes under /api/v1
	apiV1 := e.Group("/api/v1")
	routes.RouteInit(apiV1)

	e.Logger.Fatal(e.Start(":8080"))
}
