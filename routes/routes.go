package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchrule/danksongs/middleware"
)

// ErrorResponse should be used for any non-successful response
type ErrorResponse struct {
	Code    int
	Message string
}

// NewRouter creates a router with all server paths
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the DankSongs API. API can be accessed from /api"))
	})

	// See if this works
	// Song routes
	r.HandleFunc("/api/song", middleware.AuthMiddleware(CreateSongHandler)).Methods("POST")
	r.HandleFunc("/api/song", GetSongHandler).Methods("GET")
	r.HandleFunc("/api/song", middleware.AuthMiddleware(UpdateSongHandler)).Methods("PUT")
	r.HandleFunc("/api/song", middleware.AuthMiddleware(DeleteSongHandler)).Methods("DELETE")
	r.HandleFunc("/api/song/vote", middleware.AuthMiddleware(VoteHandler)).Methods("POST")

	// Playlist routes
	r.HandleFunc("/api/playlist", middleware.AuthMiddleware(CreatePlaylistHandler)).Methods("POST")
	r.HandleFunc("/api/playlist", GetPlaylistHandler).Methods("GET")
	// Delete the whole playlist
	r.HandleFunc("/api/playlist", middleware.AuthMiddleware(DeletePlaylistHandler)).Methods("DELETE")
	// Add and delete SONGS from playlist
	r.HandleFunc("/api/playlist/add", middleware.AuthMiddleware(AddSongHandler)).Methods("POST")
	r.HandleFunc("/api/playlist/remove", middleware.AuthMiddleware(RemoveSongHandler)).Methods("DELETE")

	// User routes
	r.HandleFunc("/api/user", CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/user/login", LoginUserHandler).Methods("POST")
	r.HandleFunc("/api/user/logout", middleware.AuthMiddleware(LogoutUserHandler)).Methods("POST")

	return r
}
