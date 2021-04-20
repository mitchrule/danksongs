package actions

import (
	"context"
	"log"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
)

// CreateSong adds a song to the database
func CreateSong(song models.Song) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := database.SongsCollection.InsertOne(ctx, song)

	if err != nil {
		return err
	}

	log.Println("Inserted: ", insertResult)
	return nil
}

// NewSong returns a pointer to a song with 0 votes
func newSong(title string, artist string, url string) *models.Song {
	song := models.Song{
		Title:  title,
		Artist: artist,
		URL:    url,
		Votes:  []models.Vote{},
	}

	return &song
}
