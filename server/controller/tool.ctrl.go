package controller

import "github.com/gofiber/fiber/v2"

func Tools(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"tools": "tools",
	})

}

func Tool(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"tool": "tool",
	})

}

func CreateTool(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"tool": "CreateTool",
	})

}

func RemoveTool(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"message": "Tool removed successfully",
	})

}
