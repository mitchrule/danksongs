package routes

import (
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

	// Song routes
	r.HandleFunc("/api/song", CreateSongHandler).Methods("POST")
	r.HandleFunc("/api/song", GetSongHandler).Methods("GET")
	r.HandleFunc("/api/song", middleware.AuthMiddleware(UpdateSongHandler)).Methods("PUT")
	r.HandleFunc("/api/song", middleware.AuthMiddleware(DeleteSongHandler)).Methods("DELETE")
	// r.HandleFunc("/api/song/vote/{id}", middleware.AuthMiddleware(VoteHandler)).Methods("POST")

	// Spotify API based query
	r.HandleFunc("/api/song/search", middleware.AuthMiddleware(SearchSpotifyForSongsHandler)).Methods("POST")

	// Playlist routes
	r.HandleFunc("/api/playlist", middleware.AuthMiddleware(CreatePlaylistHandler)).Methods("POST")
	r.HandleFunc("/api/playlist", GetPlaylistHandler).Methods("GET")
	// Delete the whole playlist
	r.HandleFunc("/api/playlist", middleware.AuthMiddleware(DeletePlaylistHandler)).Methods("DELETE")
	// Add and delete SONGS from playlist
	r.HandleFunc("/api/playlist/{playlistid}/add", middleware.AuthMiddleware(AddSongHandler)).Methods("POST")
	r.HandleFunc("/api/playlist/remove", middleware.AuthMiddleware(RemoveSongHandler)).Methods("DELETE")
	r.HandleFunc("/api/vote/{playlistid}/{songid}", middleware.AuthMiddleware(VoteHandler)).Methods("PUT")
	// Addional functions to retrieve playlist data
	r.HandleFunc("/api/playlist/getRecent", GetRecentPlaylistsHandler).Methods("GET")
	r.HandleFunc("/api/playlist/search", middleware.AuthMiddleware(SearchPlaylistsHandler)).Methods("POST")

	// User routes
	r.HandleFunc("/api/user", CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/user", middleware.AuthMiddleware(DeleteUserHandler)).Methods("DELETE")
	r.HandleFunc("/api/user/login", LoginUserHandler).Methods("POST")
	r.HandleFunc("/api/user/logout", middleware.AuthMiddleware(LogoutUserHandler)).Methods("POST")
	// TODO
	// r.HandleFunc("/api/user/change-password", middleware.AuthMiddleware(ChangePasswordHandler)).Methods("POST")

	return r
}
