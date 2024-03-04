package main

import (
	"github.com/EmanuelCav/alojuniordev/config"
	"github.com/EmanuelCav/alojuniordev/database"
	"github.com/EmanuelCav/alojuniordev/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	godotenv.Load()

	database.DatabaseConnection()

	app.Use(logger.New())

	routes.UserRoutes(app)
	routes.RoleRoutes(app)
	routes.ToolRoutes(app)
	routes.CategoryRoutes(app)

	err := app.Listen(":" + config.Config()["port"])

	if err != nil {
		panic("Error to connect server")
	}

}
