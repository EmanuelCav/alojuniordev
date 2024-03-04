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

func Categories(c *fiber.Ctx) error {

	var categories []models.CategoryModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := helper.ConnectionCategory().Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category models.CategoryModel

		err2 := cursor.Decode(&category)

		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err2.Error(),
			})
		}

		categories = append(categories, category)
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"categories": categories,
	})

}

func CreateCategory(c *fiber.Ctx) error {

	var category models.CategoryCreateModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.BodyParser(&category)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := helper.Validate().Struct(&category)

	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "There are empty fields",
		})
	}

	newCategory := models.CategoryModel{
		Id:        primitive.NewObjectID(),
		Category:  category.Category,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err2 := helper.ConnectionCategory().InsertOne(ctx, newCategory)

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err2.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"category": newCategory,
		"message":  "Category created successfully",
	})

}

func RemoveCategory(c *fiber.Ctx) error {

	var category models.CategoryModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	categoryId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err2 := helper.ConnectionCategory().FindOne(ctx, bson.M{"_id": categoryId}).Decode(&category)

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "User does not exists",
		})
	}

	helper.ConnectionCategory().FindOneAndDelete(ctx, bson.M{"_id": categoryId})

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"message": "Category removed successfully",
	})

}
