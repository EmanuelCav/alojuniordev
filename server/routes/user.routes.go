package routes

import (
	"github.com/EmanuelCav/alojuniordev/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controller.Users)
	app.Get("/users/:id", controller.User)
	app.Post("/users/register", controller.Register)
	app.Post("/users/login", controller.Login)
	app.Delete("/users/:id", controller.RemoveUser)
}
