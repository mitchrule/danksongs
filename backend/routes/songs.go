package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
)

// CreateSongHandler for creating new songs
func CreateSongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Values("content-type")[0] != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		res := ErrorResponse{
			Code:    400,
			Message: "Incorrect content-type",
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(payload)
	}

	var song models.Song

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &song)
	if err != nil {
		res := ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error Unmarshaling body",
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(payload)
	}

	// Placing authentication here for proof of concept

	songID, err := actions.CreateSong(song)
	if err != nil {
		res := ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}

		payload, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(payload)
	}

	songByte, err := json.Marshal(songID)

	w.WriteHeader(http.StatusOK)
	w.Write(songByte)
}

func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit GetSongHandler"))
}

func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit UpdateSongHandler"))
}

func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit DeleteSongHandler"))
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit VoteHandler"))
}
