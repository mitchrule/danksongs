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

	newPlaylist := models.Playlist{
		Name:          playListName,
		Songs:         nil,
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
func AddSong(playList models.Playlist, song models.Song) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Add the song to the array
	playList.Songs = append(playList.Songs, song)

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

	// Find Playlist
	//var playList models.Playlist
	err := database.PlaylistCollection.FindOne(ctx, bson.M{"_id": playList.ID}).Decode(&playList)

	if err != nil {
		return false, err
	}

	// remove the associated object id from the splice
	playList.Songs = removeSongFromSlice(playList.Songs, song)

	// Update playlist
	var oldPlaylist models.Playlist
	err = database.PlaylistCollection.FindOneAndUpdate(ctx, bson.M{"_id": playList.ID}, playList).Decode(&oldPlaylist)

	// Return result
	if err != nil {
		return false, err
	}

	return true, nil
}

/// Helper functions

// Removes an element from a slice because it isnt built in for no reason
// Runs in n time which probably wont work for voting
// TODO: Find a way to do this in <O(n) time
func removeSongFromSlice(songList []models.Song, songTBR models.Song) []models.Song {

	// Get a Hexadecimal representation of the Objects
	var newSongList []models.Song

	// for i := 0; i < len(songList); i++ {
	// 	if (songList[i].ID) != songTBR.ID {
	// 		append(newSongList, songList[i])
	// 	}
	// }

	// Return eveything except that index
	return newSongList
}
