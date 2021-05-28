package actions

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
	"github.com/zmb3/spotify"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2/clientcredentials"
)

// CreateSong adds a song to the database
func CreateSong(song models.Song) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := database.SongsCollection.InsertOne(ctx, song)

	if err != nil {
		return primitive.NilObjectID, err
	}

	log.Println("Inserted: ", insertResult)
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

// NewSong returns a pointer to a song with 0 votes
func newSong(title string, artist string, uri string) *models.Song {
	song := models.Song{
		Title:  title,
		Artist: artist,
		URI:    uri,
		Votes:  []models.Vote{},
	}

	return &song
}

// searchSpotifyForSong querys spotifys api for a particular song based
// on a string that the user inputs and returns an array of songs
func SearchSpotifyForSongs(query string) ([]models.Song, error) {
	log.Println("Searching...")

	// Config inorder to access Spotify API (might move to database or make an init for spotify)
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	// log.Println("Current ID")
	// log.Println(os.LookupEnv("SPOTIFY_ID"))

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	client := spotify.Authenticator{}.NewClient(token)

	// Query Spotify based on search term
	results, err := client.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	var songs []models.Song

	// handle song results and convert them into
	// an array of song structs that we can use
	if results.Tracks != nil {
		for _, item := range results.Tracks.Tracks {
			// log.Println("Other Assoicated Data: ")
			// log.Println(item)

			artistsString := ""
			for _, artists := range item.Artists {

				artistsString += artists.Name + " ,"
			}

			// Assign the associated
			// detail from the item to our struct
			currSong := models.Song{
				SpotifyID: item.ID,
				Title:     item.Name,
				Artist:    artistsString,
				URI:       string(item.URI),
				Votes:     []models.Vote{},
			}

			songs = append(songs, currSong)
		}
	}

	// Return songs or an associated error
	if err != nil {
		return nil, err
	} else {
		return songs, nil
	}
}
