package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchrule/danksongs/actions"
	"github.com/mitchrule/danksongs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// Handles the /songs/search API route
func SearchSpotifyForSongsHandler(w http.ResponseWriter, r *http.Request) {
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

	var searchTerm string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &searchTerm)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	songs, err := actions.SearchSpotifyForSongs(searchTerm)

	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	songsByte, err := json.Marshal(songs)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		// Returns the playlist as a JSON
		w.WriteHeader(http.StatusOK)
		w.Write(songsByte)
	}
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
	// Get playlist
	// Find song in playlist
	// 	Increment vote counter
	// Update playlist database record
	params := mux.Vars(r)
	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	tokenString := token.Value

	playlistID, err := primitive.ObjectIDFromHex(params["playlistid"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	songID, err := primitive.ObjectIDFromHex(params["songid"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := actions.GetUserFromToken(tokenString)
	log.Println(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	playlist, err := actions.VoteOnSong(playlistID, songID, user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(playlist)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Vote added")
	w.Write(res)
}
