package database

import (
	"github.com/EmanuelCav/alojuniordev/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(collectioName string) *mongo.Collection {

	collection := DatabaseConnection().Database(config.Config()["database"]).Collection(collectioName)

	return collection

}
