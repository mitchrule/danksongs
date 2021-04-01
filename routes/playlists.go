package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
)


func CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Values("content-type")[0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		res := ErrorResponse{
			Code: 400,
			Message: "Incorrect content-type",
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(payload)
		return
	}

	var playlist models.Playlist

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)	
	}

	err = json.Unmarshal(body, &playlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = actions.CreatePlaylist(playlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
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
