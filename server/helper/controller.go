package helper

import (
	"github.com/EmanuelCav/alojuniordev/config"
	"github.com/EmanuelCav/alojuniordev/database"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectionUser() *mongo.Collection {
	return database.GetCollection(config.Config()["userCollection"])
}

func ConnectionRole() *mongo.Collection {
	return database.GetCollection(config.Config()["roleCollection"])
}

func Validate() *validator.Validate {
	return validator.New()
}
