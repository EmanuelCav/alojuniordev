package routes

import (
	"github.com/EmanuelCav/alojuniordev/controller"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Get("/categories", controller.Categories)
	app.Post("/categories", controller.CreateCategory)
	app.Delete("/categories/:id", controller.RemoveCategory)
}
