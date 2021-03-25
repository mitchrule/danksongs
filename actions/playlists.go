package actions

import (
	"context"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatePlaylists creates an empty playlist for songs to be
// added to and returns the mongo reference if the operaion is sucessful
func CreatePlaylist(playListName string) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newPlaylist := models.Playlist{
		Name:          playListName,
		Songs:         nil,
		VoteThreshold: 0,
	}

	// TODO
	// Check for duplicates

	insertResult, err := database.SongsCollection.InsertOne(ctx, newPlaylist)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

// GetPlaylist should return the playlist specified based on the objectID specified
func GetPlaylist(playListID primitive.ObjectID) (models.Playlist, error) {
	return models.Playlist{}, nil
}

// GetPlaylist should delete the playlist specified and return a true falue
// if it suceeds and an error otherwise
func DeletePlaylist(playListID primitive.ObjectID) (bool, error) {
	return false, nil
}

// AddSong will add a song to the Songs section of a playlist by referencing
// its object id to the playlist and return a true value if successful
func AddSong(playListID primitive.ObjectID, songID primitive.ObjectID) (bool, error) {
	return false, nil
}

// RemoveSong will remove a song from a playlist based on its object id and
// return a true value if it is sucessful
func RemoveSong(playListID primitive.ObjectID, songID primitive.ObjectID) (bool, error) {
	return false, nil
}
