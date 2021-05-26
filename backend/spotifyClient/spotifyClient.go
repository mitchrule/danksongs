package spotifyClient

import (
	"context"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Instantiate a blank client
var Client = spotify.Client{}

func initSpotify() {
	// Config inorder to access Spotify API (might move to database or make an init for spotify)
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	// Generate a token that can be used within OAuth
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	// Generate a client that can be accessed
	Client = spotify.Authenticator{}.NewClient(token)
}
