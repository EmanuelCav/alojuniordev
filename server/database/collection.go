package database

import (
	"github.com/EmanuelCav/alojuniordev/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(collectioName string) *mongo.Collection {

	collection := DatabaseConnection().Database(config.Database()).Collection(collectioName)

	return collection

}
