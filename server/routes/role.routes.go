package routes

import (
	"github.com/EmanuelCav/alojuniordev/controller"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	app.Get("/roles", controller.Roles)
	app.Post("/roles", controller.CreateRole)
	app.Delete("/roles/:id", controller.RemoveRole)
}
