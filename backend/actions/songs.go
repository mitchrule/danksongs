package actions

import (
	"context"
	"log"
	"time"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/zmb3/spotify"
	// "golang.org/x/oauth2/clientcredentials"
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
func newSong(title string, artist string, url string) *models.Song {
	song := models.Song{
		Title:  title,
		Artist: artist,
		URL:    url,
		Votes:  []models.Vote{},
	}

	return &song
}

// TODO Implement this
func searchSpotifyForSongs(song string) []*models.Song {

	// Take query term

	// Query Spotify based on search term

	// Return assoicated songs as an array of songs

	// Example to use

	/*
		config := &clientcredentials.Config{
			ClientID:     os.Getenv("SPOTIFY_ID"),
			ClientSecret: os.Getenv("SPOTIFY_SECRET"),
			TokenURL:     spotify.TokenURL,
		}
		token, err := config.Token(context.Background())
		if err != nil {
			log.Fatalf("couldn't get token: %v", err)
		}

		client := spotify.Authenticator{}.NewClient(token)
		// search for playlists and albums containing "holiday"
		results, err := client.Search("holiday", spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
		if err != nil {
			log.Fatal(err)
		}

		// handle album results
		if results.Albums != nil {
			fmt.Println("Albums:")
			for _, item := range results.Albums.Albums {
				fmt.Println("   ", item.Name)
			}
		}
		// handle playlist results
		if results.Playlists != nil {
			fmt.Println("Playlists:")
			for _, item := range results.Playlists.Playlists {
				fmt.Println("   ", item.Name)
			}
		}
	*/

	return nil
}
