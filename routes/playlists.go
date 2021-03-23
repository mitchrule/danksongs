package routes

import "net/http"

func CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit CreatePlaylistHandler"))
}

func GetPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit GetPlaylistHandler"))
}

func DeletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit DeletePlaylistHandler"))
}

func AddSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit AddSongHandler"))
}

func RemoveSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit RemoveSongHandler"))
}
