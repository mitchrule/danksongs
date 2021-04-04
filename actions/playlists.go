package actions

import (
	"context"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatePlaylists creates an empty playlist for songs to be
// added to and returns the mongo reference if the operaion is sucessful
func CreatePlaylist(playListName string) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var songList = make(map[primitive.ObjectID]models.Song)

	newPlaylist := models.Playlist{
		Name:          playListName,
		Songs:         songList,
		VoteThreshold: 0,
	}

	// TODO
	// Check for duplicates

	insertResult, err := database.PlaylistCollection.InsertOne(ctx, newPlaylist)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

// GetPlaylist should return the playlist specified based on the objectID specified
func GetPlaylist(playListID primitive.ObjectID) (models.Playlist, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check mongoDB for the associated objectID
	var playList models.Playlist
	err := database.PlaylistCollection.FindOne(ctx, bson.M{"_id": playListID}).Decode(&playList)

	if err != nil {
		return models.Playlist{}, err
	} else {
		//return the playlist
		return playList, nil
	}
}

// GetPlaylist should delete the playlist specified and return a true falue
// if it suceeds and an error otherwise
func DeletePlaylist(playListID primitive.ObjectID) (bool, error) {

	// Look up the object Id
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check mongoDB for the associated objectID
	var deletedPlaylist models.Playlist
	err := database.PlaylistCollection.FindOneAndDelete(ctx, bson.M{"_id": playListID}).Decode(&deletedPlaylist)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// AddSong will add a song to the Songs section of a playlist by referencing
// its object id to the playlist and return a true value if successful
// NOTE: idk if i should use the entire model or just the Id's
func AddSong(playList models.Playlist, song models.Song) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Add the song to the map
	playList.Songs[song.ID] = song

	// Update playlist
	var oldPlaylist models.Playlist
	err := database.PlaylistCollection.FindOneAndUpdate(ctx, bson.M{"_id": playList.ID}, playList).Decode(&oldPlaylist)

	// Return result
	if err != nil {
		return false, err
	}

	return true, nil
}

// RemoveSong will remove a song from a playlist based on its object id and
// return a true value if it is sucessful
func RemoveSong(playList models.Playlist, song models.Song) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// remove the associated object id from the map
	delete(playList.Songs, song.ID)

	// Update playlist
	var newPlaylist models.Playlist
	err := database.PlaylistCollection.FindOneAndUpdate(ctx, bson.M{"_id": playList.ID}, playList).Decode(&newPlaylist)

	// Return result
	if err != nil {
		return false, err
	}

	return true, nil
}
