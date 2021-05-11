package actions

import (
	"context"
	"log"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const NUM_PLAYLISTS_RETURNED = 30

// AddSong will add a song to the Songs section of a playlist by referencing
// its object id to the playlist and return a true value if successful
// NOTE: idk if i should use the entire model or just the Id's
func AddSong(ids models.SongPLPair) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve Playlist and Song
	var playList models.Playlist
	var song models.Song

	err := database.PlaylistCollection.FindOne(ctx, bson.M{"_id": ids.PlaylistID}).Decode(&playList)

	// Error Check
	if err != nil {
		return false, err
	}

	err = database.SongsCollection.FindOne(ctx, bson.M{"_id": ids.SongID}).Decode(&song)

	// Error Check
	if err != nil {
		return false, err
	}

	// Add the song to the slice
	newSongs := append(playList.Songs, song)
	playList.Songs = newSongs

	log.Println("Current New Playlist")
	log.Println(playList)

	// Update playlist
	var oldPlaylist models.Playlist
	err = database.PlaylistCollection.FindOneAndReplace(ctx, bson.M{"_id": playList.ID}, playList).Decode(&oldPlaylist)

	// Return result
	if err != nil {
		return false, err
	}

	return true, nil
}

// CreatePlaylists creates an empty playlist for songs to be
// added to and returns the mongo reference if the operaion is sucessful
func CreatePlaylist(data models.PlaylistData) (primitive.ObjectID, error) {

	playListName := data.PlayListName
	voteThreshold := data.VoteThreshold
	votePreportion := data.VotePreportion

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newPlaylist := models.Playlist{
		Name:           playListName,
		Songs:          []models.Song{},
		VoteThreshold:  voteThreshold,
		VotePreportion: votePreportion,
	}

	// TODO
	// 1. Check for duplicates
	// 2. Add min votes required field
	// 3. Set threshold

	insertResult, err := database.PlaylistCollection.InsertOne(ctx, newPlaylist)

	if err != nil {
		return primitive.NilObjectID, err
	}

	log.Println(insertResult)

	return insertResult.InsertedID.(primitive.ObjectID), nil
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

// GetRecentPlaylists returns the most recent playlists based on how recently they
// were created. Currently set to get playlists from the last month.
func GetRecentPlaylists() ([]models.Playlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var playLists []models.Playlist

	// Sort the object ids as they are produced in an order
	// TODO: Change this to a createdAt field one day
	opts := options.Find().SetSort(bson.D{{"_id", 1}})

	// Sort for most recent
	cursor, err := database.PlaylistCollection.Find(ctx, bson.D{}, opts)

	// Return them in an array
	err = cursor.All(context.TODO(), &playLists)

	numPlaylists := len(playLists)

	// Return the result
	if err != nil {
		return nil, err
	} else {
		if numPlaylists < NUM_PLAYLISTS_RETURNED {
			return playLists, nil
		} else {
			return playLists[:NUM_PLAYLISTS_RETURNED], nil
		}
	}
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

// RemoveSong will remove a song from a playlist based on its object id and
// return a true value if it is sucessful
func RemoveSong(ids models.SongPLPair) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve Playlist and Song
	var playList models.Playlist
	var song models.Song

	err := database.PlaylistCollection.FindOne(ctx, bson.M{"_id": ids.PlaylistID}).Decode(&playList)

	// Error Check
	if err != nil {
		return false, err
	}

	err = database.SongsCollection.FindOne(ctx, bson.M{"_id": ids.SongID}).Decode(&song)

	// Error Check
	if err != nil {
		return false, err
	}

	var newSongs []models.Song

	// remove the associated object id from the map
	found := false

	for _, song := range playList.Songs {
		if song.ID != ids.SongID && !found {
			newSongs = append(newSongs, song)
			found = true
		}
	}

	if !found {
		return false, mongo.ErrNoDocuments
	}

	// Assign the new song list
	playList.Songs = newSongs

	// Update playlist
	var oldPlaylist models.Playlist
	err = database.PlaylistCollection.FindOneAndReplace(ctx, bson.M{"_id": playList.ID}, playList).Decode(&oldPlaylist)

	// Error Check
	if err != nil {
		return false, err
	}

	return true, nil
}

// SearchPlaylists will return the playlists that are most alike to
// the search term imputted by the user

// TODO finish this
func SearchPlaylists(query string) ([]models.Playlist, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var playLists []models.Playlist

	filter := bson.M{"Name": query}

	// Set the filter to apply on
	// filter := bson.D{{"hello", "world"}}

	// Sort the object ids as they are produced in an order
	opts := options.Find().SetSort(bson.D{{"_id", 1}})

	// Search for playlists based on the query
	cursor, err := database.PlaylistCollection.Find(ctx, filter, opts)

	// Error Check
	if err != nil {
		return nil, err
	}

	// Return them in an array
	err = cursor.All(context.TODO(), &playLists)

	log.Println("Current Playlists: ")
	log.Println(playLists)

	// Return the result
	if err != nil {
		return nil, err
	} else {
		return playLists, nil
	}
}

func VoteOnSong(playlistID string, songID string, username string) (models.Playlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var playlist models.Playlist

	err := database.PlaylistCollection.FindOne(ctx, bson.D{{"_id", playlistID}}).Decode(&playlist)
	if err != nil {
		log.Println("Failed to get playlist")
		return models.Playlist{}, err
	}

	for _, song := range playlist.Songs {
		if song.ID.String() == songID {

		}
	}
}
