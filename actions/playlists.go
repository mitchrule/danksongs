package actions

import (
	"context"
	"log"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
)


func CreatePlaylist(playlist models.Playlist) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := database.PlaylistsCollection.InsertOne(ctx, playlist)
	if err != nil {
		return err
	}

	log.Println("Inserted: ", insertResult)
	return nil
}