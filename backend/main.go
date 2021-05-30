package main

import (
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/database"
	"github.com/mitchrule/danksongs/routes"
	"github.com/mitchrule/danksongs/spotifyClient"
)

func main() {

	// Intantiate MongoDB and
	// Spotify API's
	database.InitDatabase()
	spotifyClient.InitSpotify()

	router := routes.NewRouter()

	s := &http.Server{
		Addr:    ":80",
		Handler: router,
	}
	log.Println("Listening on 80...")
	s.ListenAndServe()
}
