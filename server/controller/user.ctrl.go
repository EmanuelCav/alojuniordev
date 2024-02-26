package controller

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/config"
	"github.com/EmanuelCav/alojuniordev/database"
	"github.com/EmanuelCav/alojuniordev/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var connection = database.GetCollection(config.Config()["userCollection"])

func Users(c *fiber.Ctx) error {

	var users []models.UserModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := connection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err,
		})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.UserModel
		err2 := cursor.Decode(&user)

		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err2,
			})
		}

		users = append(users, user)
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"users": users,
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
