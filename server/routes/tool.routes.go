package routes

import (
	"github.com/EmanuelCav/alojuniordev/controller"
	"github.com/gofiber/fiber/v2"
)

func ToolRoutes(app *fiber.App) {
	app.Get("/tools", controller.Tools)
	app.Get("/tools/:id", controller.Tool)
	app.Post("/tools", controller.CreateTool)
	app.Delete("/tools/:id", controller.RemoveTool)
}
