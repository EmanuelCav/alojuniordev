package controller

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/config"
	"github.com/EmanuelCav/alojuniordev/helper"
	"github.com/EmanuelCav/alojuniordev/middleware"
	"github.com/EmanuelCav/alojuniordev/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Users(c *fiber.Ctx) error {

	var users []models.UserModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := helper.ConnectionUser().Find(ctx, bson.M{}, helper.UsersFilter())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.UserModel
		err2 := cursor.Decode(&user)

		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": err2.Error(),
			})
		}

		users = append(users, user)
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"users": users,
	})

}

func User(c *fiber.Ctx) error {

	var user models.UserModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err2 := helper.ConnectionUser().FindOne(ctx, bson.M{"_id": userId}, helper.UserFilter()).Decode(&user)

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "User does not exists",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user": user,
	})

}

func Register(c *fiber.Ctx) error {

	var user models.UserRegisterModel
	var roleUser models.RoleModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	validationErr := helper.Validate().Struct(&user)

	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Hay campos vacios. Por favor completa",
		})
	}

	errValidation := middleware.RegisterValid(user)

	if errValidation != "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": errValidation,
		})
	}

	var err2 error

	if user.Role == "" {
		err2 = helper.ConnectionRole().FindOne(ctx, bson.M{"role": config.Config()["userMainRole"]}).Decode(&roleUser)
	} else {
		err2 = helper.ConnectionRole().FindOne(ctx, bson.M{"role": user.Role}).Decode(&roleUser)
	}

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err2.Error(),
		})
	}

	user.Password = helper.HashPassword(user.Password)

	newUser := models.UserModel{
		Id:         primitive.NewObjectID(),
		Username:   user.Username,
		Password:   user.Password,
		Email:      user.Email,
		Role:       roleUser.Id,
		Reputation: 0,
		Status:     true,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err3 := helper.ConnectionUser().InsertOne(ctx, newUser)

	if err3 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err3.Error(),
		})
	}

	// userLoggedIn :=

	token := helper.GenerateToken(newUser.Id)

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user":    newUser,
		"token":   token,
		"message": "Bienvenido A Lo Junior!",
	})

}

func Login(c *fiber.Ctx) error {

	var user models.UserLoginModel
	var userLoggedIn models.UserModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	errValidator := helper.Validate().Struct(&user)

	if errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Hay campos vacios. Por favor completa",
		})
	}

	errValidation := middleware.LoginValid(user)

	if errValidation != "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": errValidation,
		})
	}

	helper.ConnectionUser().FindOne(ctx, bson.M{"email": user.Email}, helper.UserFilter()).Decode(&userLoggedIn)

	token := helper.GenerateToken(userLoggedIn.Id)

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"user":  userLoggedIn,
		"token": token,
	})

}

func RemoveUser(c *fiber.Ctx) error {

	var user models.UserModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	err2 := helper.ConnectionUser().FindOne(ctx, bson.M{"_id": userId}).Decode(&user)

	if err2 != nil {
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"message": "User does not exists",
		})
	}

	err3 := helper.ConnectionUser().FindOneAndDelete(ctx, bson.M{"_id": userId}).Decode(&user)

	if err3 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err3.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"message": "User removed successfully",
	})

}
