package routes

import "github.com/gofiber/fiber/v2"

func ProjectsRoutes(app *fiber.App) {
	app.Get("/projects")
	app.Get("/projects/:id")
	app.Post("/projects")
	app.Delete("/projects/:id")
}
