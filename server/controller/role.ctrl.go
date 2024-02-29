package controller

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/helper"
	"github.com/EmanuelCav/alojuniordev/middleware"
	"github.com/EmanuelCav/alojuniordev/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Roles(c *fiber.Ctx) error {

	var roles []models.RoleModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := helper.ConnectionRole().Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var role models.RoleModel
		err2 := cursor.Decode(&role)

		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err2.Error(),
			})
		}

		roles = append(roles, role)
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"roles": roles,
	})

}

func CreateRole(c *fiber.Ctx) error {

	var role models.RoleCreateModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.BodyParser(&role)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := helper.Validate().Struct(&role)

	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "There are empty fields",
		})
	}

	errValidation := middleware.RoleValid(role)

	if errValidation != "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": errValidation,
		})
	}

	newRole := models.RoleModel{
		Id:        primitive.NewObjectID(),
		Role:      role.Role,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err2 := helper.ConnectionRole().InsertOne(ctx, newRole)

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err2.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"role":    newRole,
		"message": "Role created successfully",
	})

}

func RemoveRole(c *fiber.Ctx) error {

	var role models.RoleModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	roleId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err2 := helper.ConnectionRole().FindOne(ctx, bson.M{"_id": roleId}).Decode(&role)

	if err2 != nil {
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"message": "Role does not exists",
		})
	}

	err3 := helper.ConnectionRole().FindOneAndDelete(ctx, bson.M{"_id": roleId}).Decode(&role)

	if err3 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err3.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"message": "Role removed successfully",
	})

}
