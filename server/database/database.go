package database

import (
	"context"
	"log"
	"time"

	"github.com/EmanuelCav/alojuniordev/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DatabaseConnection() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri()))

	if err != nil {
		log.Fatal(err)
	}

	err2 := client.Ping(ctx, readpref.Primary())

	if err2 != nil {
		log.Fatal(err)
	}

	log.Println("Database is running")

	return client

}
