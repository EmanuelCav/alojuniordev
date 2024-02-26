package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"users": "Users",
	})

}

func User(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user": "User",
	})

}

func Register(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user": "Register",
	})

}

func Login(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user": "Login",
	})

}

func RemoveUser(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user": "RemoveUser",
	})

}
