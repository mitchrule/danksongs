package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/mitchrule/danksongs/middleware"
)

// ErrorResponse should be used for any non-successful response
type ErrorResponse struct {
	Code    int
	Message string
}

// The path to the static directory and path to index file
// Used to serve the static React app
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(path)

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, r.URL.Path)
	fmt.Println(path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// NewRouter creates a router with all server paths
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Song routes
	r.HandleFunc("/api/song", CreateSongHandler).Methods("POST")
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

	// Serve the static React app
	spa := spaHandler{staticPath: "ui/build/", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	return r
}
