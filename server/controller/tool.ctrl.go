package controller

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/helper"
	"github.com/EmanuelCav/alojuniordev/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Tools(c *fiber.Ctx) error {

	var tools []models.ToolModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := helper.ConnectionTool().Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var tool models.ToolModel
		err2 := cursor.Decode(&tool)

		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err2.Error(),
			})
		}

		tools = append(tools, tool)
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"tools": tools,
	})

}

func Tool(c *fiber.Ctx) error {

	var tool models.ToolModel

	id := c.Params("id")

	toolId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err2 := helper.ConnectionTool().FindOne(ctx, bson.M{"_id": toolId}).Decode(&tool)

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Tool does not exists",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"tool": tool,
	})

}

func CreateTool(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"tool": "CreateTool",
	})

}

func RemoveTool(c *fiber.Ctx) error {

	var tool models.ToolModel

	id := c.Params("id")

	toolId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err2 := helper.ConnectionTool().FindOne(ctx, bson.M{"_id": toolId}).Decode(&tool)

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Tool does not exists",
		})
	}

	err3 := helper.ConnectionTool().FindOneAndDelete(ctx, bson.M{"_id": toolId}).Decode(&tool)

	if err3 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err3.Error(),
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
		"message": "Tool removed successfully",
	})

}
