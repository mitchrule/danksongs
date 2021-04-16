package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SongsCollection in MongoDB
var SongsCollection *mongo.Collection = new(mongo.Collection)

// UsersCollection in MongoDB
var UsersCollection *mongo.Collection = new(mongo.Collection)

// UsersCollection in MongoDB
var JWTCollection *mongo.Collection = new(mongo.Collection)

// Playlists
var PlaylistsCollection *mongo.Collection = new(mongo.Collection)

// InitDatabase initialises a global database client
func InitDatabase() {
	// mongoUsername := os.Getenv("MONGOUSERNAME")
	databaseName := "danksongs"
	// mongoPassword := os.Getenv("MONGOPWD")
	// mongoURI := os.Getenv("MONGO_URI")
	mongoURI := "mongodb://mongo:27017"

	// mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.sn8oj.mongodb.net/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, databaseName)
	// mongoURI := fmt.Sprintf("mongodb://%s:%s@mongo:27017/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, databaseName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	// Where the collections are initalised
	SongsCollection = client.Database(databaseName).Collection("songs")
	PlaylistsCollection = client.Database(databaseName).Collection("playlists")
	UsersCollection = client.Database(databaseName).Collection("users")
	JWTCollection = client.Database(databaseName).Collection("tokens")

	// log.Println("Database initialised", songsCollection)
}
