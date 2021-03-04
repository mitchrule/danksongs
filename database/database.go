package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SongsCollection in MongoDB
var SongsCollection *mongo.Collection = new(mongo.Collection)

// InitDatabase initialises a global database client
func InitDatabase() {
	mongoUsername := os.Getenv("MONGOUSERNAME")
	databaseName := os.Getenv("DATABASENAME")
	mongoPassword := os.Getenv("MONGOPWD")

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.sn8oj.mongodb.net/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, databaseName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	SongsCollection = client.Database(databaseName).Collection("songs")
	// log.Println("Database initialised", songsCollection)
}
