package mongo

import (
	"context"
	config "devices/lib/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cnf config.Config = config.LoadConfig()
var ctx context.Context = context.Background()

// Locates and returns the mongo collection by name
func MongoCollection(collection string) *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cnf.Mongo.Uri))

	if err != nil {
		panic(err)
	}

	result := client.Database(cnf.Mongo.Database).Collection(collection)

	return result
}
