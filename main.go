package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	mongoUsername := os.Getenv("MONGOUSERNAME")
	databaseName := os.Getenv("DATABASENAME")
	mongoPassword := os.Getenv("MONGOPWD")

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.sn8oj.mongodb.net/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, databaseName)
	fmt.Println(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	insertSampleSong(ctx, client)
}

func insertSampleSong(ctx context.Context, client *mongo.Client) {
	database := client.Database(os.Getenv("DATABASENAME"))
	songsCollection := database.Collection("songs")

	sampleSong := models.Song{
		Title:  "test song",
		Artist: "peepee",
		URL:    "www.google.com",
		Votes:  0,
	}

	insertResult, err := songsCollection.InsertOne(ctx, sampleSong)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertResult.InsertedID)
}
